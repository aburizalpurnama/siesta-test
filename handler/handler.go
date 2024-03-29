package handler

import (
	"encoding/json"
	"strconv"

	"github.com/aburizalpurnama/siesta-test/model"
	"github.com/aburizalpurnama/siesta-test/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type (
	MbtHandler interface {
		CreateSimulation(c *fiber.Ctx) error
		ApproveLending(c *fiber.Ctx) error
		SelectRepayments(c *fiber.Ctx) error
	}

	MbtHandlerImpl struct {
		mbtUsecase usecase.MbtUsecase
	}
)

func NewMbtHandler(usecase usecase.MbtUsecase) MbtHandler {
	return &MbtHandlerImpl{usecase}
}

func (h *MbtHandlerImpl) CreateSimulation(c *fiber.Ctx) error {
	accountId := c.Params("accountId")
	accId, _ := strconv.Atoi(accountId)
	var reqPayload model.CreateSimulationRequest
	var errResponse model.ErrorResponse
	err := json.Unmarshal(c.Body(), &reqPayload)
	if err != nil {
		log.Info(err)
		errResponse.Error = "invalid request body"
		c.JSON(errResponse)
		c.SendStatus(fiber.ErrBadRequest.Code)
		return err
	}

	reqPayload.AccountId = uint(accId)
	respPayload, err := h.mbtUsecase.CreateSimulation(reqPayload)
	if err != nil {
		c.SendStatus(fiber.ErrInternalServerError.Code)
		return err
	}

	c.JSON(respPayload)
	c.SendStatus(fiber.StatusOK)
	return nil
}

func (h *MbtHandlerImpl) ApproveLending(c *fiber.Ctx) error {
	lendingId := c.Params("lendingId")
	lenId, _ := strconv.Atoi(lendingId)

	reqData := model.ApproveLendingRequest{LendingId: uint(lenId)}
	err := h.mbtUsecase.ApproveLending(reqData)
	if err != nil {
		c.SendStatus(fiber.ErrInternalServerError.Code)
		return err
	}

	c.SendStatus(fiber.StatusOK)
	return nil
}

func (h *MbtHandlerImpl) SelectRepayments(c *fiber.Ctx) error {
	accountId := c.Params("accountId")
	accId, _ := strconv.Atoi(accountId)

	reqData := model.SelectRepaymentsRequest{AccountId: uint(accId)}
	respData, err := h.mbtUsecase.SelectRepayments(reqData)
	if err != nil {
		c.SendStatus(fiber.ErrInternalServerError.Code)
		return err
	}

	c.JSON(respData)
	c.SendStatus(fiber.StatusOK)
	return nil
}
