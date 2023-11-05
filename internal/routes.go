package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func setup(c *gin.Context) {

}

func getUserById(c *gin.Context) {
	userid := c.Param("id")
	id, _ := strconv.Atoi(userid)
	row := DB.QueryRow(fmt.Sprintf("select * from users where user_id=%s", id))
	var user User
	err := row.Scan(&user.UserId, &user.Name, &user.Nickname, &user.Rate, &user.Description,
		&user.Login, &user.Password)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		user.Password = ""
		c.JSON(200, gin.H{
			"user": user,
		})
	}
}

func getTagsByUser(c *gin.Context) {
	ids := c.Param("id")
	//id, _ := strconv.Atoi(ids)
	rows, err := DB.Query(fmt.Sprintf("SELECT tag_id FROM users_tags WHERE user_id=%s", ids))
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		tags := []Tag{}
		for rows.Next() {
			var tag Tag
			rows.Scan(&tag.TagID, &tag.GlobalTagID, &tag.Activity)
			tags = append(tags, tag)
		}
		c.JSON(200, gin.H{
			"tags": tags,
		})
	}
}

func getAllTags(c *gin.Context) {
	tags := []Tag{}
	rows, err := DB.Query("SELECT * FROM tag")
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		for rows.Next() {
			var tagg Tag
			_ = rows.Scan(&tagg.TagID, &tagg.Activity, &tagg.GlobalTagID)
			tags = append(tags, tagg)
		}
		c.JSON(200, gin.H{
			"tags": tags,
		})
	}
}

func getTeamsByUser(c *gin.Context) {
	ids := c.Param("id")
	teams := []Team{}
	rows, err := DB.Query(fmt.Sprintf("SELECT team_id FROM user_team WHERE user_id=%s", ids))
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		for rows.Next() {
			var team Team
			rows.Scan(&team.TeamID, &team.Name, &team.Rate, &team.Description, &team.Rules, &team.RegDate, &team.Place)
			teams = append(teams, team)
		}
		c.JSON(200, gin.H{
			"teams": teams,
		})
	}
}

func getTeamById(c *gin.Context) {
	ids := c.Param("id")
	var team Team
	row := DB.QueryRow(fmt.Sprintf("SELECT * from teams where team_id=%s", ids))
	err := row.Scan(&team.TeamID, &team.Name, &team.Rate, &team.Description, &team.Rules, &team.RegDate, &team.Place)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		c.JSON(200, gin.H{
			"team": team,
		})
	}
}

func getTagsByTeam(c *gin.Context) {
	ids := c.Param("id")
	tags := []Tag{}
	rows, err := DB.Query(fmt.Sprintf("select tag_id from teams_tags where tag_id=%s", ids))
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		for rows.Next() {
			var tag Tag
			rows.Scan(&tag.TagID, &tag.GlobalTagID, &tag.Activity)
			tags = append(tags, tag)
		}
		c.JSON(200, gin.H{
			"tags": tags,
		})
	}
}
