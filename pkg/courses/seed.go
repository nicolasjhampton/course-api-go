package courses

func seedCourses() {
	DB.DropTable(&Course{})
	DB.AutoMigrate(&Course{})
	DB.Create(&Course{
		Title: "Build a Basic Bookcase",
	})
	DB.Create(&Course{
		Title: "Learn How to Program",
	})
	DB.Create(&Course{
		Title: "New Course Updated Again Hello",
	})
}



// CourseDetails{
// 	Id: "57029ed4795118be119cc43d",
// 	User: User{
// 		Id:       "57029ed4795118be119cc437",
// 		Name:     "Joe Smith",
// 		Email:    "joe@smith.com",
// 		Password: "$2a$10$8k/BgPdYO88b0yQ/oQ9ij.l4K/mpmzNRJHbhjdPLgcvlhjsP4FJEq",
// 		Version:  "0",
// 	},
// 	Title: "Build a Basic Bookcase",
// 	Description: `High-end furniture projects are great to dream 
// 	about. But unless you have a well-equipped shop and some 
// 	serious woodworking experience to draw on, it can be difficult 
// 	to turn the dream into a reality.\n\nNot every piece of 
// 	furniture needs to be a museum showpiece, though.`,
// 	Time: "12 hours",
// 	Materials: `* 1/2 x 3/4 inch parting strip\n* 1 x 2 common pine\n
// 	* 1 x 4 common pine\n* 1 x 10 common pine\n* 1/4 inch thick lauan plywood\n
// 	* Finishing Nails\n* Sandpaper\n* Wood Glue\n* Wood Filler\n
// 	* Minwax Oil Based Polyurethane\n`,
// 	Version: "0",
// 	Reviews: []Review{
// 		Review{
// 			Id:     "57029ed4795118be119cc43a",
// 			UserId: "57029ed4795118be119cc439",
// 			Rating: "5",
// 			Review: `Lorem ipsum dolor sit amet, consectetur adipisicing 
// 			elit, sed do eiusmod tempor incididunt ut labore et dolore 
// 			magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation 
// 			ullamco laboris nisi ut aliquip ex ea commodo consequat.`,
// 			Version: "0",
// 		},
// 		Review{
// 			Id:     "57029ed4795118be119cc43b",
// 			UserId: "57029ed4795118be119cc438",
// 			Rating: "3",
// 			Review: `Duis aute irure dolor in reprehenderit in voluptate 
// 			velit esse cillum dolore eu fugiat nulla pariatur. Excepteur 
// 			sint occaecat cupidatat non proident, sunt in culpa qui officia 
// 			deserunt mollit anim id est laborum.`,
// 			Version: "0",
// 		},
// 	},
// 	Steps: []Step{
// 		Step{
// 			Id:          "57029ed4795118be119cc43f",
// 			Number:      "1",
// 			Title:       "Yakkety smackity...",
// 			Description: "Here's some things...",
// 		},
// 		Step{
// 			Id:          "57029ed4795118be119cc43e",
// 			Number:      "2",
// 			Title:       "Blah Blah Blah",
// 			Description: "And some other stuff...",
// 		},
// 	},
// }
