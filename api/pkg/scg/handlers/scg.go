package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulski/test-scg/pkg/scg/adapter"
	"github.com/soulski/test-scg/pkg/scg/repository"
)

type SCGHandler struct {
	placeAdapter     *adapter.PlaceAdapter
	userRepository   repository.UserRepository
	converRepository repository.ConversationRepository
}

func NewSCGHandler(
	placeAdapter *adapter.PlaceAdapter,
	userRepository repository.UserRepository,
	converRepository repository.ConversationRepository,
) *SCGHandler {
	return &SCGHandler{placeAdapter, userRepository, converRepository}
}

type XYZQuery struct {
	Numbers []int `form:"numbers"`
}

func (h *SCGHandler) ListUsers(c *gin.Context) {
	c.JSON(http.StatusOK, h.userRepository.FindAll())
}

func (h *SCGHandler) ListConversations(c *gin.Context) {
	c.JSON(http.StatusOK, h.converRepository.FindAll())
}

func (h *SCGHandler) FindRestaurants(c *gin.Context) {
	result, err := h.placeAdapter.FindNearbyPlace()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *SCGHandler) FindXYZ(c *gin.Context) {
	const UNKNOW = -1

	var query *XYZQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	numbers := append(append([]int{UNKNOW}, query.Numbers...), UNKNOW, UNKNOW)
	diffs := make([]int, len(numbers))
	step := UNKNOW

	// find diff
	for index := range diffs {
		if numbers[index] == UNKNOW || numbers[index+1] == UNKNOW {
			diffs[index] = UNKNOW
		} else {
			diffs[index] = numbers[index+1] - numbers[index]
		}
	}

	// find step
	for index := range diffs {
		if diffs[index] != UNKNOW && diffs[index+1] != UNKNOW {
			step = diffs[index+1] - diffs[index]
			break
		}
	}

	// fill diff
	for index := len(diffs) / 2; index >= 0; index-- {
		if diffs[index] != UNKNOW {
			continue
		}

		diffs[index] = diffs[index+1] - step
	}

	for index := len(diffs) / 2; index < len(diffs); index++ {
		if diffs[index] != UNKNOW {
			continue
		}

		diffs[index] = diffs[index-1] + step
	}

	//fill numbers
	for index := len(numbers) / 2; index >= 0; index-- {
		if numbers[index] != UNKNOW {
			continue
		}

		numbers[index] = numbers[index+1] - diffs[index]
	}

	for index := len(numbers) / 2; index < len(numbers); index++ {
		if numbers[index] != UNKNOW {
			continue
		}

		numbers[index] = numbers[index-1] + diffs[index-1]
	}

	c.JSON(http.StatusOK, numbers)
}
