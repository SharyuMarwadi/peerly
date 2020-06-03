const log4js = require("log4js");

const db = require("../../models/sequelize");
const jsonwebtoken = require("../../jwtTokenValidation/jwtValidation");
const userBlacklistedTokens = db.user_blacklisted_tokens;
require("../../config/loggerConfig");

const logger = log4js.getLogger();

module.exports.logout = async (req, res) => {
  const authHeader = req.headers["authorization"];
  const token = authHeader && authHeader.split(" ")[1];
  let decode = await jsonwebtoken.getData(authHeader);
  const user = {
    user_id: decode.userId,
    token: token,
    expires_at: decode.exp,
  };
  userBlacklistedTokens
    .create(user)
    .then(() => {
      logger.info("user logout");
      logger.info("=========================================");
      res.status(200).send();
    })
    .catch(() => {
      logger.error("internal server error");
      logger.info("=========================================");
      res.status(500).send({
        error: {
          message: "internal server error",
        },
      });
    });
};
