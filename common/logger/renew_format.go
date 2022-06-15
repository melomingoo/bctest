package logger

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
)

//CodeLineNumberRenewHook 로그가 찍히는 코드라인 위치 출력을 위한 Hook
type CodeLineNumberRenewHook struct{}

//Levels CodeLineNumberRenewHook이 적용되는 로그레벨: 현재 전부, 추후 Error나 Debug시에만 출력하려면 변경 필요
func (h *CodeLineNumberRenewHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

//Fire CodeLineNumberRenewHook 구현 메소드
func (h *CodeLineNumberRenewHook) Fire(entry *logrus.Entry) error {
	// WithFields 여부에 관계 없이 정확한 위치가 출력되도록 처리
	// https://github.com/sirupsen/logrus/issues/63 참고

	pc := make([]uintptr, 3, 3)
	cnt := runtime.Callers(9, pc)

	for i := 0; i < cnt; i++ {
		fu := runtime.FuncForPC(pc[i] - 1)
		funcName := fu.Name()
		if !strings.Contains(funcName, "mobility.repo/maas/dealer/eu.service.fms.git/common/logger") {
			file, line := fu.FileLine(pc[i] - 1)
			entry.Data[LogFunc] = fmt.Sprintf("%s:%d:%s", path.Base(file), line, path.Ext(funcName)[1:])
			break
		}
	}

	return nil
}

//MaaSReNewLogFormatter MaaSLoggerFormatter 인스턴스
type MaaSReNewLogFormatter struct {
	//Hostname 로깅에 사용할 호스트 이름
	hostname string

	color bool

	once sync.Once
}

func (f *MaaSReNewLogFormatter) init() error {
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}
	f.hostname = hostname
	return nil
}

func checkIfTerminal(w io.Writer) bool {
	switch v := w.(type) {
	case *os.File:
		return terminal.IsTerminal(int(v.Fd()))
	default:
		return false
	}
}

//Format 로그 포맷팅 수행
func (f *MaaSReNewLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	f.once.Do(func() {
		f.color = checkIfTerminal(entry.Logger.Out) && (runtime.GOOS != "windows")
		f.init()
	})

	isColored := f.isColored()

	b.WriteString(entry.Time.Format("2006-01-02 15:04:05.999"))
	b.WriteByte(' ')

	level := strings.ToUpper(entry.Level.String())
	if isColored {
		f.appendValueWithColor(entry, b, entry.Data[LogRuleType], level[0:4])
	} else {
		f.appendValue(b, entry.Data[LogRuleType], level)
	}
	b.WriteByte(' ')

	b.WriteString(f.hostname)

	b.WriteByte(' ')

	b.WriteString(serviceName)

	b.WriteByte(' ')

	b.WriteString(brandName)

	b.WriteByte(' ')
	f.appendValue(b, entry.Data[LogTID], "null")
	b.WriteByte(' ')
	f.appendValue(b, entry.Data[LogFunc], "null")
	b.WriteString(" --- ")

	b.WriteString("[")
	f.appendValue(b, entry.Data[LogProtocol], "null")
	b.WriteByte(',')
	f.appendValue(b, entry.Data[LogMethod], "null")
	b.WriteByte(',')
	f.appendValue(b, entry.Data[LogURL], "null")
	b.WriteByte(',')
	f.appendValue(b, entry.Data[LogStatus], "null")
	b.WriteString("]")

	b.WriteString(" [")

	f.appendValue(b, entry.Data[LogSID], "null")
	b.WriteString("] --- ")

	if isColored {
		f.appendValueWithColor(entry, b, entry.Message, level)
	} else {
		f.appendValue(b, entry.Message, level)
	}

	b.WriteByte('\n')

	return b.Bytes(), nil
}

func (f *MaaSReNewLogFormatter) appendValue(b *bytes.Buffer, value interface{}, nilVal string) {
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

func (f *MaaSReNewLogFormatter) isColored() bool {
	return f.color
}

func (f *MaaSReNewLogFormatter) appendValueWithColor(entry *logrus.Entry, b *bytes.Buffer, value interface{}, nilVal string) {
	var levelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		levelColor = 37 // gray
	case logrus.WarnLevel:
		levelColor = 33 // yellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		levelColor = 31 // red
	default:
		levelColor = 36 // blue
	}
	if value == nil {
		b.Write([]byte(fmt.Sprintf("\x1b[%dm%s\x1b[0m", levelColor, nilVal)))
	} else {
		stringVal, ok := value.(string)
		if !ok {
			stringVal = fmt.Sprint(value)
		}
		b.Write([]byte(fmt.Sprintf("\x1b[%dm%s\x1b[0m", levelColor, stringVal)))
	}
}
