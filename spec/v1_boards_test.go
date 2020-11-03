package spec_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/keitam0/agility/spec"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("/v1/boards", func() {
	Describe("GET", func() {
		Context("with an agility instance", func() {
			var (
				ins         *os.Process
				environment map[string]string = map[string]string{}
			)

			JustBeforeEach(func() {
				i, err := spec.StartAgility(environment)
				if err != nil {
					panic(err)
				}
				ins = i
			})

			AfterEach(func() {
				ins.Kill()
				ins.Wait()
			})

			Context("some teams are given", func() {
				BeforeEach(func() {
					environment["TEAM_BOARD_IDS"] = "team-a:1000,team-b:2000"
				})

				Context("some sprints are available at JIRA", func() {
					BeforeEach(func() {
						jira := gin.New()
						jira.GET("/agile/1.0/board/:bid/sprint", func(ctx *gin.Context) {
							ctx.String(http.StatusOK, `{
								"startAt": 0,
								"isLast": true,
								"maxResults": 3,
								"values": [
									{
										"name": "p1",
										"state": "closed"
									},
									{
										"name": "p2",
										"state": "closed"
									},
									{
										"name": "p3",
										"state": "closed"
									}
								]
							}`)
						})
						jira.GET("/agile/1.0/board/:bid/issue", func(ctx *gin.Context) {
							ctx.String(http.StatusOK, `{
								"startAt": 0,
								"maxResults": 1,
								"issues": [
									{
										"status": {
											"name": "解決済み"
										},
										"customfield_10002": 5,
										"closedSprints": [
											{
												"name": "p3"
											}
										]
									}
								]
							}`)
						})
						environment["JIRA_API_ENDPOINT"] = httptest.NewServer(jira).URL
					})

					Context("gets boards", func() {
						var boards []interface{}

						JustBeforeEach(func() {
							res, err := http.Get("http://localhost:9090/v1/boards")
							if err != nil {
								panic(err)
							}
							defer res.Body.Close()
							var b []interface{}
							if err := json.NewDecoder(res.Body).Decode(&b); err != nil {
								panic(err)
							}
							boards = b
						})

						Specify("got two team boards", func() {
							Expect(boards).To(HaveLen(2))
							Expect(boards[0].(map[string]interface{})["team"]).To(Equal("team-a"))
							Expect(boards[1].(map[string]interface{})["team"]).To(Equal("team-b"))
						})

						Specify("the boards have sprints", func() {
							Expect(boards[0].(map[string]interface{})["sprints"]).To(HaveLen(3))
						})
					})
				})
			})
		})
	})
})
