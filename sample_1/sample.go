package sample_1

import (
	log "github.com/cihub/seelog"
)

func LogExample() {
	testConfig := `
<seelog type="sync">
	<outputs formatid="main">
		<console/>
	</outputs>
	<formats>
		<format id="main" format="%Date/%Time [%Level] %RelFile:%Line %Msg%n"/>
	</formats>
</seelog>`

	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.ReplaceLogger(logger)

	log.Info("Hello world!")

}
