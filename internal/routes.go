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
	//id, _ := strconv.Atoi(userid)
	row := DB.QueryRow(fmt.Sprintf("select * from users where user_id=%s", userid))
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
		c.JSON(200, user)
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
			rows.Scan(&tag.TagID)
			row := DB.QueryRow(fmt.Sprintf("select * from tag where tag_id=%s", strconv.Itoa(tag.TagID)))
			row.Scan(&tag.TagID, &tag.Activity, &tag.GlobalTagID)
			tags = append(tags, tag)
		}
		c.JSON(200, tags)
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
		c.JSON(200, tags)
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
			rows.Scan(&team.TeamID)
			row := DB.QueryRow(fmt.Sprintf("select * from team where team_id=%s", strconv.Itoa(team.TeamID)))
			row.Scan(&team.TeamID, &team.Name, &team.Rate, &team.Description, &team.Rules, &team.RegDate, &team.Place)
			teams = append(teams, team)
		}
		c.JSON(200, teams)
	}
}

func getTeamById(c *gin.Context) {
	ids := c.Param("id")
	var team Team
	row := DB.QueryRow(fmt.Sprintf("SELECT * from team where team_id=%s", ids))
	err := row.Scan(&team.TeamID, &team.Name, &team.Rate, &team.Description, &team.Rules, &team.RegDate, &team.Place)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		c.JSON(200, team)
	}
}

func getTagsByTeam(c *gin.Context) {
	ids := c.Param("id")
	tags := []Tag{}
	rows, err := DB.Query(fmt.Sprintf("select tag_id from team_tags where team_id=%s", ids))
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		for rows.Next() {
			var tag Tag
			rows.Scan(&tag.TagID)
			row := DB.QueryRow(fmt.Sprintf("select * from tag where tag_id=%s", strconv.Itoa(tag.TagID)))
			row.Scan(&tag.TagID, &tag.Activity, &tag.GlobalTagID)
			tags = append(tags, tag)
		}
		c.JSON(200, tags)
	}
}

func getUsersByTeam(c *gin.Context) {
	id := c.Param("id")
	rows, err := DB.Query(fmt.Sprintf("select user_id from user_team where team_id=%s", id))
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		usrs := []User{}
		for rows.Next() {
			var usr User
			rows.Scan(&usr.UserId)
			row := DB.QueryRow(fmt.Sprintf("select * from users where user_id=%s", strconv.Itoa(usr.UserId)))
			row.Scan(&usr.UserId, &usr.Name, &usr.Nickname, &usr.Rate, &usr.Description, &usr.Login, &usr.Password)
			usr.Password = ""
			usrs = append(usrs, usr)
		}
		c.JSON(200, usrs)
	}
}

func getEventById(c *gin.Context) {
	id := c.Param("id")
	row := DB.QueryRow(fmt.Sprintf("select * from events where event_id=%s", id))
	var event Event
	err := row.Scan(&event.EventID, &event.Name, &event.Description, &event.Date, &event.Online, &event.MainTheme, &event.Url, &event.CreatorID)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		c.JSON(200, event)
	}
}

func getTeamsByEvent(c *gin.Context) {
	id := c.Param("id")
	rows, err := DB.Query(fmt.Sprintf("select team_id from team_event where event_id=%s", id))
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		teams := []Team{}
		for rows.Next() {
			var team Team
			rows.Scan(&team.TeamID)
			row := DB.QueryRow(fmt.Sprintf("select * from team where team_id=%s", strconv.Itoa(team.TeamID)))
			row.Scan(&team.TeamID, &team.Name, &team.Rate, &team.Description, &team.Rules, &team.RegDate, &team.Place)
			teams = append(teams, team)
		}
		c.JSON(200, teams)
	}
}

func getTagsByEvent(c *gin.Context) {
	id := c.Param("id")
	rows, err := DB.Query(fmt.Sprintf("select tag_id from events_tags where event_id=%s", id))
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		tags := []Tag{}
		for rows.Next() {
			var tag Tag
			rows.Scan(&tag.TagID)
			row := DB.QueryRow(fmt.Sprintf("select * from tag where tag_id=%s", strconv.Itoa(tag.TagID)))
			row.Scan(&tag.TagID, &tag.Activity, &tag.GlobalTagID)
			tags = append(tags, tag)
		}
		c.JSON(200, tags)
	}
}

func getAllEvents(c *gin.Context) {
	events := []Event{}
	rows, err := DB.Query(fmt.Sprintf("select * from events"))
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		for rows.Next() {
			var event Event
			rows.Scan(&event.EventID, &event.Name, &event.Description, &event.Date, &event.Online, &event.MainTheme, &event.Url, &event.CreatorID)
			events = append(events, event)
		}
		c.JSON(200, events)
	}
}

func getAllTeams(c *gin.Context) {
	teams := []Team{}
	rows, err := DB.Query(fmt.Sprintf("select * from team"))
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
		c.JSON(200, teams)
	}
}

func getAllGlobalTags(c *gin.Context) {
	gtags := []GlobalTag{}
	rows, err := DB.Query("select * from global_tag")
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		for rows.Next() {
			var gtag GlobalTag
			rows.Scan(&gtag.GlobalTagID, &gtag.Category)
			gtags = append(gtags, gtag)
		}
		c.JSON(200, gtags)
	}
}

func getTagsByGlobalTag(c *gin.Context) {
	id := c.Param("id")
	rows, err := DB.Query(fmt.Sprintf("select * from tag where global_tag_id=%s", id))
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		tags := []Tag{}
		for rows.Next() {
			var tag Tag
			rows.Scan(&tag.TagID, &tag.Activity, &tag.GlobalTagID)
			tags = append(tags, tag)
		}
		c.JSON(200, tags)
	}
}

func getAllUsers(c *gin.Context) {
	rows, err := DB.Query("select * from users")
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"server": -1,
		})
	} else {
		users := []User{}
		for rows.Next() {
			var usr User
			rows.Scan(&usr.UserId, &usr.Name, &usr.Nickname, &usr.Rate, &usr.Description, &usr.Login, &usr.Password)
			usr.Password = ""
			users = append(users, usr)
		}
		c.JSON(200, users)
	}
}
