package victor

import (
	"fmt"
	"github.com/brettbuddin/victor/util/google"
	"math/rand"
	"strconv"
	"time"
)
func randomString(strSlice []string) string {
	if len(strSlice) > 0 {
		return strSlice[rand.Intn(len(strSlice))]
	}
	return "Fitstar!"
}

func defaults(robot *Robot) {
	robot.Respond("ping", func(m Message) {
		m.Reply("pong!")
	})

	robot.Respond("roll( (\\d+))?", func(m Message) {
		defer recover()

		bound := 100
		val := m.Params()[1]

		if val != "" {
			var err error
			bound, err = strconv.Atoi(val)

			if err != nil {
				return
			}
		}

		rand.Seed(time.Now().UTC().UnixNano())
		random := rand.Intn(bound)
		m.Reply(fmt.Sprintf("rolled a %d of %d", random, bound))
	})

	robot.Respond("(image|img|gif|animate) (.*)", func(m Message) {
		gifOnly := (m.Params()[0] == "gif" || m.Params()[0] == "animate")

		imgQuery := m.Params()[1]
		if imgQuery == "yonkers" || imgQuery == "Yonkers" {
			m.Room().Say(randomString([]string{
				"http://www.weinberg.northwestern.edu/alumni/crosscurrents/2012-2013-fall-winter/images/big-data/shih.jpg",
				"https://lh6.googleusercontent.com/-LTZDe1okbYM/TirGDmEkR-I/AAAAAAAAD6k/YUck4Xe62K0/w1095-h821-no/DSCN0223.JPG",
				"https://vandelaylabs1.campfirenow.com/room/562152/uploads/5250668/yup.jpg",
				"https://lh5.googleusercontent.com/-E2GIy-qEK-Y/TirD5SlrIpI/AAAAAAAAIQY/tg7hcoGmdN4/w462-h821-no/SAM_2014.JPG",
				"https://lh3.googleusercontent.com/-05n6L7MwM1M/TirC-DK5lXI/AAAAAAAAIQo/O7qgwj9XnG8/w1095-h821-no/P1020833.JPG",
				"https://dl.dropboxusercontent.com/sh/crs7ynybt64niu0/0K5gqhCPRx/Harrison/IMG_3922.JPG?token_hash=AAFcBPNuMRKbtnAE7j8sP6I18s1rckPAce2KBUZPH7ji4g",
				"https://dl.dropboxusercontent.com/sh/crs7ynybt64niu0/tZi3QrzrEG/Harrison/IMG_3928.JPG?token_hash=AAFcBPNuMRKbtnAE7j8sP6I18s1rckPAce2KBUZPH7ji4g",

				}))
			return
		}
		result, err := google.ImageSearch(m.Params()[1], gifOnly)

		if err != nil {
			m.Room().Say(fmt.Sprintf("There was error making the request: %v", err))
			return
		}

		if result == "" {
			m.Room().Say("I didn't find anything.")
			return
		}

		m.Room().Say(result)
	})
}
