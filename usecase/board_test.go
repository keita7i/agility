package usecase_test

import (
	"github.com/keitam0/agility/domain/agile"
	"github.com/keitam0/agility/usecase"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Board", func() {
	var (
		board           *usecase.Board
		jiraServiceMock *usecase.JIRAServiceMock
		teams           []string
	)

	BeforeEach(func() {
		jiraServiceMock = &usecase.JIRAServiceMock{}
	})

	JustBeforeEach(func() {
		board = &usecase.Board{
			JIRAService: jiraServiceMock,
			Teams:       teams,
		}
	})

	Describe("AllBoards()", func() {
		Context("when there are some teams", func() {
			BeforeEach(func() {
				teams = []string{"team0", "team1"}
			})

			It("responses team boards", func() {
				By("jira provides boards")
				jiraServiceMock.BoardOfTeamFunc = func(team string, maxSprints int) (agile.Board, error) {
					return agile.NewBoard(team), nil
				}
				b, err := board.AllBoards()
				Expect(err).NotTo(BeNil())
				Expect(b).To(HaveLen(2))
			})
		})
	})
})
