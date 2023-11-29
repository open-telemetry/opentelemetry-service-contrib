// Code generated by "go.opentelemetry.io/collector/cmd/builder". DO NOT EDIT.

package main

import (
	"go.opentelemetry.io/collector/connector"
	forwardconnector "go.opentelemetry.io/collector/connector/forwardconnector"
	"go.opentelemetry.io/collector/exporter"
	debugexporter "go.opentelemetry.io/collector/exporter/debugexporter"
	loggingexporter "go.opentelemetry.io/collector/exporter/loggingexporter"
	otlpexporter "go.opentelemetry.io/collector/exporter/otlpexporter"
	otlphttpexporter "go.opentelemetry.io/collector/exporter/otlphttpexporter"
	"go.opentelemetry.io/collector/extension"
	ballastextension "go.opentelemetry.io/collector/extension/ballastextension"
	zpagesextension "go.opentelemetry.io/collector/extension/zpagesextension"
	"go.opentelemetry.io/collector/otelcol"
	"go.opentelemetry.io/collector/processor"
	batchprocessor "go.opentelemetry.io/collector/processor/batchprocessor"
	memorylimiterprocessor "go.opentelemetry.io/collector/processor/memorylimiterprocessor"
	"go.opentelemetry.io/collector/receiver"
	otlpreceiver "go.opentelemetry.io/collector/receiver/otlpreceiver"

	countconnector "github.com/open-telemetry/opentelemetry-collector-contrib/connector/countconnector"
	datadogconnector "github.com/open-telemetry/opentelemetry-collector-contrib/connector/datadogconnector"
	exceptionsconnector "github.com/open-telemetry/opentelemetry-collector-contrib/connector/exceptionsconnector"
	routingconnector "github.com/open-telemetry/opentelemetry-collector-contrib/connector/routingconnector"
	servicegraphconnector "github.com/open-telemetry/opentelemetry-collector-contrib/connector/servicegraphconnector"
	spanmetricsconnector "github.com/open-telemetry/opentelemetry-collector-contrib/connector/spanmetricsconnector"
	alibabacloudlogserviceexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/alibabacloudlogserviceexporter"
	awscloudwatchlogsexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awscloudwatchlogsexporter"
	awsemfexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsemfexporter"
	awskinesisexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awskinesisexporter"
	awss3exporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awss3exporter"
	awsxrayexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/awsxrayexporter"
	azuredataexplorerexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuredataexplorerexporter"
	azuremonitorexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/azuremonitorexporter"
	carbonexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/carbonexporter"
	cassandraexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/cassandraexporter"
	clickhouseexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/clickhouseexporter"
	coralogixexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/coralogixexporter"
	datadogexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datadogexporter"
	datasetexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/datasetexporter"
	dynatraceexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/dynatraceexporter"
	elasticsearchexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter"
	f5cloudexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/f5cloudexporter"
	fileexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/fileexporter"
	googlecloudexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudexporter"
	googlecloudpubsubexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudpubsubexporter"
	googlemanagedprometheusexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlemanagedprometheusexporter"
	honeycombmarkerexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/honeycombmarkerexporter"
	influxdbexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/influxdbexporter"
	instanaexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/instanaexporter"
	kafkaexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/kafkaexporter"
	loadbalancingexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/loadbalancingexporter"
	logicmonitorexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/logicmonitorexporter"
	logzioexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/logzioexporter"
	lokiexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/lokiexporter"
	mezmoexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/mezmoexporter"
	opencensusexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/opencensusexporter"
	opensearchexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/opensearchexporter"
	prometheusexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusexporter"
	prometheusremotewriteexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/prometheusremotewriteexporter"
	pulsarexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/pulsarexporter"
	sapmexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sapmexporter"
	sentryexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sentryexporter"
	signalfxexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter"
	skywalkingexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/skywalkingexporter"
	splunkhecexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/splunkhecexporter"
	sumologicexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/sumologicexporter"
	syslogexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/syslogexporter"
	tanzuobservabilityexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/tanzuobservabilityexporter"
	tencentcloudlogserviceexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/tencentcloudlogserviceexporter"
	zipkinexporter "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/zipkinexporter"
	asapauthextension "github.com/open-telemetry/opentelemetry-collector-contrib/extension/asapauthextension"
	awsproxy "github.com/open-telemetry/opentelemetry-collector-contrib/extension/awsproxy"
	basicauthextension "github.com/open-telemetry/opentelemetry-collector-contrib/extension/basicauthextension"
	bearertokenauthextension "github.com/open-telemetry/opentelemetry-collector-contrib/extension/bearertokenauthextension"
	headerssetterextension "github.com/open-telemetry/opentelemetry-collector-contrib/extension/headerssetterextension"
	healthcheckextension "github.com/open-telemetry/opentelemetry-collector-contrib/extension/healthcheckextension"
	httpforwarder "github.com/open-telemetry/opentelemetry-collector-contrib/extension/httpforwarder"
	jaegerremotesampling "github.com/open-telemetry/opentelemetry-collector-contrib/extension/jaegerremotesampling"
	oauth2clientauthextension "github.com/open-telemetry/opentelemetry-collector-contrib/extension/oauth2clientauthextension"
	dockerobserver "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/dockerobserver"
	ecsobserver "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecsobserver"
	ecstaskobserver "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/ecstaskobserver"
	hostobserver "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/hostobserver"
	k8sobserver "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/k8sobserver"
	oidcauthextension "github.com/open-telemetry/opentelemetry-collector-contrib/extension/oidcauthextension"
	pprofextension "github.com/open-telemetry/opentelemetry-collector-contrib/extension/pprofextension"
	sigv4authextension "github.com/open-telemetry/opentelemetry-collector-contrib/extension/sigv4authextension"
	dbstorage "github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/dbstorage"
	filestorage "github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/filestorage"
	attributesprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/attributesprocessor"
	cumulativetodeltaprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/cumulativetodeltaprocessor"
	datadogprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/datadogprocessor"
	deltatorateprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/deltatorateprocessor"
	filterprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor"
	groupbyattrsprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbyattrsprocessor"
	groupbytraceprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/groupbytraceprocessor"
	k8sattributesprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/k8sattributesprocessor"
	metricsgenerationprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricsgenerationprocessor"
	metricstransformprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/metricstransformprocessor"
	probabilisticsamplerprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/probabilisticsamplerprocessor"
	redactionprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/redactionprocessor"
	remotetapprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/remotetapprocessor"
	resourcedetectionprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor"
	resourceprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourceprocessor"
	routingprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/routingprocessor"
	servicegraphprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/servicegraphprocessor"
	spanmetricsprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/spanmetricsprocessor"
	spanprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/spanprocessor"
	sumologicprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/sumologicprocessor"
	tailsamplingprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor"
	transformprocessor "github.com/open-telemetry/opentelemetry-collector-contrib/processor/transformprocessor"
	activedirectorydsreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/activedirectorydsreceiver"
	aerospikereceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/aerospikereceiver"
	apachereceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/apachereceiver"
	apachesparkreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/apachesparkreceiver"
	awscloudwatchreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscloudwatchreceiver"
	awscontainerinsightreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awscontainerinsightreceiver"
	awsecscontainermetricsreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsecscontainermetricsreceiver"
	awsfirehosereceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsfirehosereceiver"
	awsxrayreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/awsxrayreceiver"
	azureblobreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureblobreceiver"
	azureeventhubreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azureeventhubreceiver"
	azuremonitorreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azuremonitorreceiver"
	bigipreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/bigipreceiver"
	carbonreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver"
	chronyreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/chronyreceiver"
	cloudflarereceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/cloudflarereceiver"
	cloudfoundryreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/cloudfoundryreceiver"
	collectdreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/collectdreceiver"
	couchdbreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/couchdbreceiver"
	datadogreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/datadogreceiver"
	dockerstatsreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/dockerstatsreceiver"
	elasticsearchreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/elasticsearchreceiver"
	expvarreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/expvarreceiver"
	filelogreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/filelogreceiver"
	filereceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/filereceiver"
	filestatsreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/filestatsreceiver"
	flinkmetricsreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/flinkmetricsreceiver"
	fluentforwardreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/fluentforwardreceiver"
	googlecloudpubsubreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudpubsubreceiver"
	googlecloudspannerreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver"
	haproxyreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/haproxyreceiver"
	hostmetricsreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver"
	httpcheckreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/httpcheckreceiver"
	iisreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/iisreceiver"
	influxdbreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/influxdbreceiver"
	jaegerreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jaegerreceiver"
	jmxreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/jmxreceiver"
	journaldreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/journaldreceiver"
	k8sclusterreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sclusterreceiver"
	k8seventsreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8seventsreceiver"
	k8sobjectsreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/k8sobjectsreceiver"
	kafkametricsreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkametricsreceiver"
	kafkareceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkareceiver"
	kubeletstatsreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kubeletstatsreceiver"
	lokireceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/lokireceiver"
	memcachedreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/memcachedreceiver"
	mongodbatlasreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbatlasreceiver"
	mongodbreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mongodbreceiver"
	mysqlreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/mysqlreceiver"
	nginxreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nginxreceiver"
	nsxtreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/nsxtreceiver"
	opencensusreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/opencensusreceiver"
	oracledbreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/oracledbreceiver"
	otlpjsonfilereceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/otlpjsonfilereceiver"
	podmanreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/podmanreceiver"
	postgresqlreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/postgresqlreceiver"
	prometheusreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver"
	pulsarreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/pulsarreceiver"
	purefareceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/purefareceiver"
	purefbreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/purefbreceiver"
	rabbitmqreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/rabbitmqreceiver"
	receivercreator "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/receivercreator"
	redisreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/redisreceiver"
	riakreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/riakreceiver"
	sapmreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sapmreceiver"
	signalfxreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/signalfxreceiver"
	simpleprometheusreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/simpleprometheusreceiver"
	skywalkingreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/skywalkingreceiver"
	snmpreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/snmpreceiver"
	snowflakereceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/snowflakereceiver"
	solacereceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/solacereceiver"
	splunkhecreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/splunkhecreceiver"
	sqlqueryreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlqueryreceiver"
	sqlserverreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sqlserverreceiver"
	sshcheckreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/sshcheckreceiver"
	statsdreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver"
	syslogreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/syslogreceiver"
	tcplogreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/tcplogreceiver"
	udplogreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/udplogreceiver"
	vcenterreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/vcenterreceiver"
	wavefrontreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/wavefrontreceiver"
	webhookeventreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/webhookeventreceiver"
	windowseventlogreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowseventlogreceiver"
	windowsperfcountersreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowsperfcountersreceiver"
	zipkinreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zipkinreceiver"
	zookeeperreceiver "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/zookeeperreceiver"
)

func components() (otelcol.Factories, error) {
	var err error
	factories := otelcol.Factories{}

	factories.Extensions, err = extension.MakeFactoryMap(
		zpagesextension.NewFactory(),
		ballastextension.NewFactory(),
		asapauthextension.NewFactory(),
		awsproxy.NewFactory(),
		basicauthextension.NewFactory(),
		bearertokenauthextension.NewFactory(),
		headerssetterextension.NewFactory(),
		healthcheckextension.NewFactory(),
		httpforwarder.NewFactory(),
		jaegerremotesampling.NewFactory(),
		oauth2clientauthextension.NewFactory(),
		ecsobserver.NewFactory(),
		ecstaskobserver.NewFactory(),
		hostobserver.NewFactory(),
		k8sobserver.NewFactory(),
		dockerobserver.NewFactory(),
		oidcauthextension.NewFactory(),
		pprofextension.NewFactory(),
		sigv4authextension.NewFactory(),
		filestorage.NewFactory(),
		dbstorage.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}

	factories.Receivers, err = receiver.MakeFactoryMap(
		otlpreceiver.NewFactory(),
		activedirectorydsreceiver.NewFactory(),
		aerospikereceiver.NewFactory(),
		apachereceiver.NewFactory(),
		apachesparkreceiver.NewFactory(),
		awscloudwatchreceiver.NewFactory(),
		awscontainerinsightreceiver.NewFactory(),
		awsecscontainermetricsreceiver.NewFactory(),
		awsfirehosereceiver.NewFactory(),
		awsxrayreceiver.NewFactory(),
		azureblobreceiver.NewFactory(),
		azureeventhubreceiver.NewFactory(),
		azuremonitorreceiver.NewFactory(),
		bigipreceiver.NewFactory(),
		carbonreceiver.NewFactory(),
		chronyreceiver.NewFactory(),
		cloudflarereceiver.NewFactory(),
		cloudfoundryreceiver.NewFactory(),
		collectdreceiver.NewFactory(),
		couchdbreceiver.NewFactory(),
		datadogreceiver.NewFactory(),
		dockerstatsreceiver.NewFactory(),
		elasticsearchreceiver.NewFactory(),
		expvarreceiver.NewFactory(),
		filelogreceiver.NewFactory(),
		filestatsreceiver.NewFactory(),
		filereceiver.NewFactory(),
		flinkmetricsreceiver.NewFactory(),
		fluentforwardreceiver.NewFactory(),
		googlecloudpubsubreceiver.NewFactory(),
		googlecloudspannerreceiver.NewFactory(),
		haproxyreceiver.NewFactory(),
		hostmetricsreceiver.NewFactory(),
		httpcheckreceiver.NewFactory(),
		influxdbreceiver.NewFactory(),
		iisreceiver.NewFactory(),
		jaegerreceiver.NewFactory(),
		jmxreceiver.NewFactory(),
		journaldreceiver.NewFactory(),
		k8sclusterreceiver.NewFactory(),
		k8seventsreceiver.NewFactory(),
		k8sobjectsreceiver.NewFactory(),
		kafkametricsreceiver.NewFactory(),
		kafkareceiver.NewFactory(),
		kubeletstatsreceiver.NewFactory(),
		lokireceiver.NewFactory(),
		memcachedreceiver.NewFactory(),
		mongodbatlasreceiver.NewFactory(),
		mongodbreceiver.NewFactory(),
		mysqlreceiver.NewFactory(),
		nginxreceiver.NewFactory(),
		nsxtreceiver.NewFactory(),
		opencensusreceiver.NewFactory(),
		oracledbreceiver.NewFactory(),
		otlpjsonfilereceiver.NewFactory(),
		podmanreceiver.NewFactory(),
		postgresqlreceiver.NewFactory(),
		prometheusreceiver.NewFactory(),
		pulsarreceiver.NewFactory(),
		purefareceiver.NewFactory(),
		purefbreceiver.NewFactory(),
		rabbitmqreceiver.NewFactory(),
		receivercreator.NewFactory(),
		redisreceiver.NewFactory(),
		riakreceiver.NewFactory(),
		sapmreceiver.NewFactory(),
		signalfxreceiver.NewFactory(),
		simpleprometheusreceiver.NewFactory(),
		skywalkingreceiver.NewFactory(),
		snowflakereceiver.NewFactory(),
		solacereceiver.NewFactory(),
		splunkhecreceiver.NewFactory(),
		sqlqueryreceiver.NewFactory(),
		sqlserverreceiver.NewFactory(),
		sshcheckreceiver.NewFactory(),
		statsdreceiver.NewFactory(),
		syslogreceiver.NewFactory(),
		tcplogreceiver.NewFactory(),
		udplogreceiver.NewFactory(),
		vcenterreceiver.NewFactory(),
		wavefrontreceiver.NewFactory(),
		webhookeventreceiver.NewFactory(),
		snmpreceiver.NewFactory(),
		windowsperfcountersreceiver.NewFactory(),
		windowseventlogreceiver.NewFactory(),
		zipkinreceiver.NewFactory(),
		zookeeperreceiver.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}

	factories.Exporters, err = exporter.MakeFactoryMap(
		debugexporter.NewFactory(),
		loggingexporter.NewFactory(),
		otlpexporter.NewFactory(),
		otlphttpexporter.NewFactory(),
		alibabacloudlogserviceexporter.NewFactory(),
		awscloudwatchlogsexporter.NewFactory(),
		awsemfexporter.NewFactory(),
		awskinesisexporter.NewFactory(),
		awss3exporter.NewFactory(),
		awsxrayexporter.NewFactory(),
		azuredataexplorerexporter.NewFactory(),
		azuremonitorexporter.NewFactory(),
		carbonexporter.NewFactory(),
		clickhouseexporter.NewFactory(),
		cassandraexporter.NewFactory(),
		coralogixexporter.NewFactory(),
		datadogexporter.NewFactory(),
		datasetexporter.NewFactory(),
		dynatraceexporter.NewFactory(),
		elasticsearchexporter.NewFactory(),
		f5cloudexporter.NewFactory(),
		fileexporter.NewFactory(),
		googlecloudexporter.NewFactory(),
		googlecloudpubsubexporter.NewFactory(),
		googlemanagedprometheusexporter.NewFactory(),
		honeycombmarkerexporter.NewFactory(),
		influxdbexporter.NewFactory(),
		instanaexporter.NewFactory(),
		kafkaexporter.NewFactory(),
		loadbalancingexporter.NewFactory(),
		logicmonitorexporter.NewFactory(),
		logzioexporter.NewFactory(),
		lokiexporter.NewFactory(),
		mezmoexporter.NewFactory(),
		opencensusexporter.NewFactory(),
		opensearchexporter.NewFactory(),
		prometheusexporter.NewFactory(),
		prometheusremotewriteexporter.NewFactory(),
		pulsarexporter.NewFactory(),
		sapmexporter.NewFactory(),
		sentryexporter.NewFactory(),
		signalfxexporter.NewFactory(),
		skywalkingexporter.NewFactory(),
		splunkhecexporter.NewFactory(),
		sumologicexporter.NewFactory(),
		syslogexporter.NewFactory(),
		tanzuobservabilityexporter.NewFactory(),
		tencentcloudlogserviceexporter.NewFactory(),
		zipkinexporter.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}

	factories.Processors, err = processor.MakeFactoryMap(
		batchprocessor.NewFactory(),
		memorylimiterprocessor.NewFactory(),
		attributesprocessor.NewFactory(),
		cumulativetodeltaprocessor.NewFactory(),
		datadogprocessor.NewFactory(),
		deltatorateprocessor.NewFactory(),
		filterprocessor.NewFactory(),
		groupbyattrsprocessor.NewFactory(),
		groupbytraceprocessor.NewFactory(),
		k8sattributesprocessor.NewFactory(),
		metricsgenerationprocessor.NewFactory(),
		metricstransformprocessor.NewFactory(),
		probabilisticsamplerprocessor.NewFactory(),
		redactionprocessor.NewFactory(),
		resourcedetectionprocessor.NewFactory(),
		resourceprocessor.NewFactory(),
		routingprocessor.NewFactory(),
		servicegraphprocessor.NewFactory(),
		spanmetricsprocessor.NewFactory(),
		sumologicprocessor.NewFactory(),
		spanprocessor.NewFactory(),
		tailsamplingprocessor.NewFactory(),
		transformprocessor.NewFactory(),
		remotetapprocessor.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}

	factories.Connectors, err = connector.MakeFactoryMap(
		forwardconnector.NewFactory(),
		countconnector.NewFactory(),
		datadogconnector.NewFactory(),
		exceptionsconnector.NewFactory(),
		routingconnector.NewFactory(),
		servicegraphconnector.NewFactory(),
		spanmetricsconnector.NewFactory(),
	)
	if err != nil {
		return otelcol.Factories{}, err
	}

	return factories, nil
}
