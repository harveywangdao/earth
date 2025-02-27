var log4js = require("log4js");
//var logger = log4js.getLogger();
//logger.level = "debug";
//logger.debug("Some debug messages");

log4js.configure({
    appenders: { // 输出源
      out: { type: "stdout" },
      app: { type: "file", filename: "application.log" },
      err: { type: 'stderr' }
    },
    categories: { // 类别
      default: { appenders: ["out", "app"], level: "debug" },
      normal: { appenders: ["out", "app"], level: "info" },
      err: { appenders: ["err"], level: "error" }
    },
})

//const logger = log4js.getLogger("normal")
var logger = log4js.getLogger('nodeJS');

logger.info("我是信息日志")
logger.error("我是错误日志")
