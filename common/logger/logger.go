package logger

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	//LogTID Request ID
	LogTID = "tid"

	//LogSID Span ID
	LogSID = "sid"

	//LogFunc 로깅이 일어나는 함수, 라인, 파일 등 정보, CodeLineNumberHook을 사용해서 설정해야 함
	LogFunc = "func"

	//LogRuleType RuleType으로 Request, Response 등의 값으로 설정, 설정하지 않으면 로그 레벨
	LogRuleType = "ruletype"

	//LogRuleRequest LogRuleType 중 Request 상수값
	LogRuleRequest = "Request"

	//LogRuleResponse LogRuleType 중 Response 상수값
	LogRuleResponse = "Response"

	//LogProtocol GRPC, HTTP, MQTT 등 프로토콜
	LogProtocol = "protocol"

	//LogMethod GET, PUT, POST, DELETE 등
	LogMethod = "method"

	//LogURL http 엔드포인트
	LogURL = "url"

	//LogStatus 호출 결과 코드
	LogStatus = "status"
)

//CodeLineNumberHook 로그가 찍히는 코드라인 위치 출력을 위한 Hook
type CodeLineNumberHook struct{}

//Levels CodeLineNumberHook이 적용되는 로그레벨: 현재 전부, 추후 Error나 Debug시에만 출력하려면 변경 필요
func (h *CodeLineNumberHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

//Fire CodeLineNumberHook 구현 메소드
func (h *CodeLineNumberHook) Fire(entry *logrus.Entry) error {
	// WithFields 여부에 관계 없이 정확한 위치가 출력되도록 처리
	// https://github.com/sirupsen/logrus/issues/63 참고

	pc := make([]uintptr, 3, 3)
	cnt := runtime.Callers(6, pc)

	for i := 0; i < cnt; i++ {
		fu := runtime.FuncForPC(pc[i] - 1)
		funcName := fu.Name()

		if !strings.Contains(funcName, "github.com/sirupsen/logrus") {
			file, line := fu.FileLine(pc[i] - 1)

			entry.Data[LogFunc] = fmt.Sprintf("%s:%d:%s", path.Base(file), line, path.Ext(funcName)[1:])

			break
		}
	}

	return nil
}

type LogFormatter struct {
	//Hostname 로깅에 사용할 호스트 이름
	Hostname string

	//Servername 서비스 명으로 설정
	Servername string

	//Brand 서비스 브랜드
	Brand string
}

//Init LogFormatter 초기화 함수, hostname을 초기화
func (f *LogFormatter) Init() error {
	if f.Hostname != "" {
		return nil
	}

	hostname, err := os.Hostname()

	if err != nil {
		return err
	}

	f.Hostname = hostname

	return nil
}

//Format 로그 포맷팅 수행
func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	b.WriteString(entry.Time.Format("2006-01-02 15:04:05.999"))
	b.WriteByte(' ')

	level := strings.ToUpper(entry.Level.String())
	b.WriteString(level)
	b.WriteByte(' ')

	if f.Hostname == "" {
		hostname, err := os.Hostname()

		if err != nil {
			return nil, err
		}

		b.WriteString(hostname)
	} else {
		b.WriteString(f.Hostname)
	}

	b.WriteByte(' ')

	if f.Servername == "" {
		b.WriteString("Unknown")
	} else {
		b.WriteString(f.Servername)
	}

	b.WriteByte(' ')

	if f.Brand == "" {
		b.WriteString("null")
	} else {
		b.WriteString(f.Brand)
	}

	b.WriteByte(' ')

	f.appendValue(b, entry.Data[LogTID], "null")

	b.WriteByte(' ')

	f.appendValue(b, entry.Data[LogFunc], "null")

	b.WriteString(" --- ")

	f.appendValue(b, entry.Data[LogRuleType], level)

	b.WriteString(" [")

	f.appendValue(b, entry.Data[LogProtocol], "null")

	b.WriteByte(',')

	f.appendValue(b, entry.Data[LogMethod], "null")

	b.WriteByte(',')

	f.appendValue(b, entry.Data[LogURL], "null")

	b.WriteByte(',')

	f.appendValue(b, entry.Data[LogStatus], "null")

	b.WriteString("] [")

	f.appendValue(b, entry.Data[LogSID], "null")

	b.WriteString("] --- ")

	b.WriteString(entry.Message)

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func (f *LogFormatter) appendValue(b *bytes.Buffer, value interface{}, nilVal string) {
	if value == nil {
		b.WriteString(nilVal)
	} else {
		stringVal, ok := value.(string)
		if !ok {
			stringVal = fmt.Sprint(value)
		}

		b.WriteString(stringVal)
	}
}

func InitServiceLogger(serviceName string, brandName string, protocol string, debug bool) *logrus.Entry {
	formatter := &LogFormatter{
		Servername: serviceName,
		Brand:      brandName,
	}

	// Init 호출로 호스트 이름 설정 등 초기화
	formatter.Init()

	if err := formatter.Init(); err != nil {
		logrus.Errorf("formatter.Init() error = %v", err)
	}

	logrus.SetFormatter(formatter)
	logrus.AddHook(new(CodeLineNumberHook))

	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	logger := logrus.WithField(LogProtocol, protocol)

	return logger
}
