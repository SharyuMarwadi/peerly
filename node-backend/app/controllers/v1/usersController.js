const qs = require("qs"); //eslint-disable-line node/no-extraneous-require
const moment = require("moment");
const log4js = require("log4js");

const utility = require("../../utils/utility");
const db = require("../../models/sequelize");
const jwtToken = require("../../jwtTokenValidation/jwtValidation");
const validateSchema = require("./validationSchema/UsersValidationSchema");
const constant = require("../../constant/resConstants");
require("../../config/loggerConfig");

const logger = log4js.getLogger();
const Users = db.users;

module.exports.findUsersByOrg = async (req, res) => {
  const authHeader = req.headers["authorization"];
  let userData = await jwtToken.getData(authHeader);
  const schema = validateSchema.findUsersByOrg();
  let obj = qs.parse(req.query);
  let limitOffsetObj = await utility.getLimitAndOffset(obj);

  schema
    .validate(obj, { abortEarly: false })
    .then(async () => {
      if (obj.org_id) {
        let superAdminAuth = await utility.validateRole(userData.roleId, [
          "SuperAdmin",
        ]);
        if (!superAdminAuth) {
          {
            logger.error("find user by orrganisaiton access denied");
            logger.info("user id: " + userData.userId);
            logger.info("=========================================");
            res.status(403).send({
              error: {
                code: "access_denied",
                message: "Permission required",
              },
            });
          }
        } else {
          Users.findAll({
            where: { org_id: obj.org_id },
            attributes: { exclude: ["soft_delete"] },
            order: [["id", "ASC"]],
            limit: limitOffsetObj.limit,
            offset: limitOffsetObj.offset,
          })
            .then((users) => {
              res.status(200).send({
                data: users,
              });
            })
            .catch(() => {
              logger.error("Error executing find users by organisation");
              logger.info("user id: " + userData.userId);
              logger.error(constant.INTRENAL_SERVER_ERROR_MESSAGE);
              logger.info("=========================================");
              res
                .status(500)
                .send(
                  utility.getErrorResponseObject(
                    constant.INTRENAL_SERVER_ERROR_CODE,
                    constant.INTRENAL_SERVER_ERROR_MESSAGE
                  )
                );
            });
        }
      } else {
        Users.findAll({
          where: { org_id: userData.orgId, role_id: 3 },
          attributes: { exclude: ["soft_delete"] },
          order: [["id", "ASC"]],
          limit: limitOffsetObj.limit,
          offset: limitOffsetObj.offset,
        })
          .then((users) => {
            res.status(200).send({
              data: users,
            });
          })
          .catch(() => {
            logger.error("Error executing find users by organisation");
            logger.info("user id: " + userData.userId);
            logger.error(constant.INTRENAL_SERVER_ERROR_MESSAGE);
            logger.info("=========================================");
            res
              .status(500)
              .send(
                utility.getErrorResponseObject(
                  constant.INTRENAL_SERVER_ERROR_CODE,
                  constant.INTRENAL_SERVER_ERROR_MESSAGE
                )
              );
          });
      }
    })
    .catch((err) => {
      logger.error("validation error");
      logger.error(JSON.stringify(err));
      logger.info("=========================================");
      res.status(400).send({
        error: utility.getFormattedErrorObj(
          constant.INVALID_QUERY_PARAMS_CODE,
          constant.INVALID_QUERY_PARAMS_MESSAGE,
          err.errors
        ),
      });
    });
};

module.exports.getProfile = async (req, res) => {
  let userData = await jwtToken.getData(req.headers["authorization"]);

  Users.findByPk(userData.userId, {
    attributes: {
      exclude: ["soft_delete", "soft_delete_by", "soft_delete_at"],
    },
  })
    .then((profile) => {
      res.status(200).send({
        data: profile,
      });
    })
    .catch(() => {
      logger.error("executing getProfile");
      logger.info("user id: " + userData.userId);
      logger.error(constant.INTRENAL_SERVER_ERROR_MESSAGE);
      logger.info("=========================================");
      res
        .status(500)
        .send(
          utility.getErrorResponseObject(
            constant.INTRENAL_SERVER_ERROR_CODE,
            constant.INTRENAL_SERVER_ERROR_MESSAGE
          )
        );
    });
};

module.exports.getProfileById = async (req, res) => {
  let userData = await jwtToken.getData(req.headers["authorization"]);
  const id = req.params.id;
  const idSchema = validateSchema.getProfileById();

  idSchema
    .validate({ id }, { abortEarly: false })
    .then(() => {
      Users.findAll({
        where: { id: id },
        attributes: {
          exclude: ["soft_delete", "soft_delete_by", "soft_delete_at"],
        },
      })
        .then((data) => {
          if (data.length != 0) {
            res.status(200).send({
              data: data,
            });
          } else {
            logger.error(constant.USER_NOT_FOUND_MESSAGE);
            logger.info("user id: " + userData.userId);
            logger.info("=========================================");
            res
              .status(404)
              .send(
                utility.getErrorResponseObject(
                  constant.USER_NOT_FOUND_CODE,
                  constant.USER_NOT_FOUND_MESSAGE
                )
              );
          }
        })
        .catch(() => {
          logger.error("Error in getProfileById");
          logger.info("user id: " + userData.userId);
          logger.error(constant.INTRENAL_SERVER_ERROR_MESSAGE);
          logger.info("=========================================");
          res
            .status(500)
            .send(
              utility.getErrorResponseObject(
                constant.INTRENAL_SERVER_ERROR_CODE,
                constant.INTRENAL_SERVER_ERROR_MESSAGE
              )
            );
        });
    })
    .catch((err) => {
      logger.error("validation error");
      logger.error(JSON.stringify(err));
      logger.info("=========================================");
      res.status(400).send({
        error: utility.getFormattedErrorObj(
          constant.INVALID_USER_CODE,
          constant.INVALID_USER_MESSAGE,
          err.errors
        ),
      });
    });
};

module.exports.updateUser = async (req, res) => {
  let userData = await jwtToken.getData(req.headers["authorization"]);
  const schema = validateSchema.updateUser();
  const updateUser = {
    first_name: req.body.first_name,
    last_name: req.body.last_name,
    display_name: req.body.display_name,
    profile_image_url: req.body.profile_image_url,
  };
  logger.info("executing update user");
  logger.info("user id: " + userData.userId);
  logger.info(JSON.stringify(updateUser));
  logger.info("=========================================");

  schema
    .validate(updateUser, { abortEarly: false })
    .then(() => {
      Users.update(updateUser, {
        returning: true,
        where: { id: userData.userId },
      })
        .then(([rowsUpdate, [updateUsers]]) => {
          if (rowsUpdate == 1) {
            res.status(200).send({
              data: updateUsers,
            });
          } else {
            logger.error(constant.USER_NOT_FOUND_MESSAGE);
            logger.info("user id: " + userData.userId);
            logger.info("=========================================");
            res
              .status(404)
              .send(
                utility.getErrorResponseObject(
                  constant.USER_NOT_FOUND_CODE,
                  constant.USER_NOT_FOUND_MESSAGE
                )
              );
          }
        })
        .catch(() => {
          logger.error("Error in updating user");
          logger.info("user id: " + userData.userId);
          logger.error(constant.INTRENAL_SERVER_ERROR_MESSAGE);
          logger.info("=========================================");
          res
            .status(500)
            .send(
              utility.getErrorResponseObject(
                constant.INTRENAL_SERVER_ERROR_CODE,
                constant.INTRENAL_SERVER_ERROR_MESSAGE
              )
            );
        });
    })
    .catch((err) => {
      logger.error("validation error");
      logger.error(JSON.stringify(err));
      logger.info("=========================================");
      res.status(400).send({
        error: utility.getFormattedErrorObj(
          constant.INVALID_USER_CODE,
          constant.INVALID_USER_MESSAGE,
          err.errors
        ),
      });
    });
};

module.exports.updateUserByAdmin = async (req, res) => {
  let userData = await jwtToken.getData(req.headers["authorization"]);
  const schema = validateSchema.updateUserByAdmin();
  const updateUserByAdmin = {
    role_id: req.body.role_id,
  };
  const updateUserValidation = {
    role_id: req.body.role_id,
    userId: req.params.id,
  };
  logger.info("executing update user by admin");
  logger.info("user id: " + userData.userId);
  logger.info(JSON.stringify(updateUserValidation));
  logger.info("=========================================");

  schema
    .validate(updateUserValidation, { abortEarly: false })
    .then(() => {
      Users.update(updateUserByAdmin, {
        returning: true,
        where: { id: req.params.id },
      })
        .then(([rowsUpdate, [updateUserByAdmin]]) => {
          if (rowsUpdate == 1) {
            res.status(200).send({
              data: updateUserByAdmin,
            });
          } else {
            logger.error(constant.USER_NOT_FOUND_MESSAGE);
            logger.info("user id: " + userData.userId);
            logger.info("=========================================");
            res
              .status(404)
              .send(
                utility.getErrorResponseObject(
                  constant.USER_NOT_FOUND_CODE,
                  constant.USER_NOT_FOUND_MESSAGE
                )
              );
          }
        })
        .catch(() => {
          logger.error("Error in updating user");
          logger.info("user id: " + userData.userId);
          logger.error(constant.INTRENAL_SERVER_ERROR_MESSAGE);
          logger.info("=========================================");
          res
            .status(500)
            .send(
              utility.getErrorResponseObject(
                constant.INTRENAL_SERVER_ERROR_CODE,
                constant.INTRENAL_SERVER_ERROR_MESSAGE
              )
            );
        });
    })
    .catch((err) => {
      logger.error("validation error");
      logger.error(JSON.stringify(err));
      logger.info("=========================================");
      res.status(400).send({
        error: utility.getFormattedErrorObj(
          constant.INVALID_USER_CODE,
          constant.INVALID_USER_MESSAGE,
          err.errors
        ),
      });
    });
};

module.exports.deleteUser = async (req, res) => {
  const authHeader = req.headers["authorization"];
  let userData = await jwtToken.getData(authHeader);
  const schema = validateSchema.softDeleteUser();
  const softDeleteUser = {
    soft_delete: true,
    soft_delete_by: userData.userId,
    soft_delete_at: moment.utc().unix(),
  };
  logger.info("executing soft delete user");
  logger.info("user id: " + userData.userId);
  logger.info(JSON.stringify(softDeleteUser));
  logger.info("=========================================");

  schema
    .validate({ id: req.params.id }, { abortEarly: false })
    .then(() => {
      Users.update(softDeleteUser, {
        returning: true,
        where: { id: req.params.id },
      })
        .then(([rowsDelete]) => {
          if (rowsDelete) {
            res.status(200).send();
          } else {
            logger.error("error at executing soft delete user");
            logger.info("user id: " + userData.userId);
            logger.error(constant.USER_NOT_FOUND_MESSAGE + req.params.id);
            logger.info("=========================================");
            res
              .status(404)
              .send(
                utility.getErrorResponseObject(
                  constant.USER_NOT_FOUND_CODE,
                  constant.USER_NOT_FOUND_MESSAGE
                )
              );
          }
        })
        .catch(() => {
          logger.error("error at executing soft delete user");
          logger.error("user id: " + userData.userId);
          logger.error(constant.INTRENAL_SERVER_ERROR_MESSAGE);
          logger.info("=========================================");
          res
            .status(500)
            .send(
              utility.getErrorResponseObject(
                constant.INTRENAL_SERVER_ERROR_CODE,
                constant.INTRENAL_SERVER_ERROR_MESSAGE
              )
            );
        });
    })
    .catch((err) => {
      logger.error("validation error at soft delete user");
      logger.info("user id: " + userData.userId);
      logger.error(JSON.stringify(err));
      logger.info("=========================================");
      res.status(400).send({
        error: utility.getFormattedErrorObj(
          constant.INVALID_USER_CODE,
          constant.INVALID_USER_MESSAGE,
          err.errors
        ),
      });
    });
};
