const log4js = require("log4js");

log4js.addLayout("json", function (config) {
  console.log(config);
  return function (logEvent) {
    console.log(logEvent);
    return JSON.stringify(logEvent) + config.separator;
  };
});

log4js.configure({
  appenders: {
    out: { type: "stdout", layout: { type: "json", separator: "," } },
    //out: { type: "stdout"},
  },
  categories: {
    default: { appenders: ["out"], level: "info" },
  },
});

const logger = log4js.getLogger("json-test");
logger.info("this is just a test");
logger.error("of a custom appender");
logger.warn("that outputs json");
log4js.shutdown(() => {});
