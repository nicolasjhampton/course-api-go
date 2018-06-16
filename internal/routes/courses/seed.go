package courses

import (
	m "github.com/nicolasjhampton/course-api-go/internal/models"
)

func seedCourses() {
	DB.DropTable(&m.Course{})
	DB.AutoMigrate(&m.Course{})
	DB.Create(&m.Course{
		UserID: 1,
		Title:  "Build a Basic Bookcase",
		Description: `High-end furniture projects are great to dream 
		about. But unless you have a well-equipped shop and some 
		serious woodworking experience to draw on, it can be difficult 
		to turn the dream into a reality.\n\nNot every piece of 
		furniture needs to be a museum showpiece, though.`,
		Time: "12 hours",
		Materials: `* 1/2 x 3/4 inch parting strip\n* 1 x 2 common pine\n
		* 1 x 4 common pine\n* 1 x 10 common pine\n* 1/4 inch thick lauan plywood\n
		* Finishing Nails\n* Sandpaper\n* Wood Glue\n* Wood Filler\n
		* Minwax Oil Based Polyurethane\n`,
		Version: "0",
	})
	DB.Create(&m.Course{
		UserID:      2,
		Title:       "Learn How to Program",
		Description: "In this course, you'll learn how to write code like a pro!",
		Time:        "6 hours",
		Materials:   "* Notebook computer running Mac OS X or Windows\n* Text editor",
		Version:     "0",
	})
	// DB.Create(&m.Course{
	// 	Title: "New Course Updated Again Hello",
	// })
}

// &m.Course{
// 	UserID: 2,
// 	Title: "Learn How to Program",
// 	Description: "In this course, you'll learn how to write code like a pro!",
// 	Time: "6 hours",
// 	Materials: "* Notebook computer running Mac OS X or Windows\n* Text editor",
// 	Version: "0",
// }

// Reviews: []Review{
// 	Review{
// 		Id:     "57029ed4795118be119cc43a",
// 		UserId: "57029ed4795118be119cc439",
// 		Rating: "5",
// 		Review: `Lorem ipsum dolor sit amet, consectetur adipisicing
// 		elit, sed do eiusmod tempor incididunt ut labore et dolore
// 		magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation
// 		ullamco laboris nisi ut aliquip ex ea commodo consequat.`,
// 		Version: "0",
// 	},
// 	Review{
// 		Id:     "57029ed4795118be119cc43b",
// 		UserId: "57029ed4795118be119cc438",
// 		Rating: "3",
// 		Review: `Duis aute irure dolor in reprehenderit in voluptate
// 		velit esse cillum dolore eu fugiat nulla pariatur. Excepteur
// 		sint occaecat cupidatat non proident, sunt in culpa qui officia
// 		deserunt mollit anim id est laborum.`,
// 		Version: "0",
// 	},
// },
// Steps: []Step{
// 	Step{
// 		Id:          "57029ed4795118be119cc43f",
// 		Number:      "1",
// 		Title:       "Yakkety smackity...",
// 		Description: "Here's some things...",
// 	},
// 	Step{
// 		Id:          "57029ed4795118be119cc43e",
// 		Number:      "2",
// 		Title:       "Blah Blah Blah",
// 		Description: "And some other stuff...",
// 	},
// },
// }

// "one": {
// 	"_id": "57029ed4795118be119cc43d",
// 	"user": "->users.joesmith",
// 	"title": "Build a Basic Bookcase",
// 	"description": "High-end furniture projects are great to dream about. But unless you have a well-equipped shop and some serious woodworking experience to draw on, it can be difficult to turn the dream into a reality.\n\nNot every piece of furniture needs to be a museum showpiece, though. Often a simple design does the job just as well and the experience gained in completing it goes a long way toward making the next project even better.\n\nOur pine bookcase, for example, features simple construction and it's designed to be built with basic woodworking tools. Yet, the finished project is a worthy and useful addition to any room of the house. While it's meant to rest on the floor, you can convert the bookcase to a wall-mounted storage unit by leaving off the baseboard. You can secure the cabinet to the wall by screwing through the cabinet cleats into the wall studs.\n\nWe made the case out of materials available at most building-supply dealers and lumberyards, including 1/2 x 3/4-in. parting strip, 1 x 2, 1 x 4 and 1 x 10 common pine and 1/4-in.-thick lauan plywood. Assembly is quick and easy with glue and nails, and when you're done with construction you have the option of a painted or clear finish.\n\nAs for basic tools, you'll need a portable circular saw, hammer, block plane, combination square, tape measure, metal rule, two clamps, nail set and putty knife. Other supplies include glue, nails, sandpaper, wood filler and varnish or paint and shellac.\n\nThe specifications that follow will produce a bookcase with overall dimensions of 10 3/4 in. deep x 34 in. wide x 48 in. tall. While the depth of the case is directly tied to the 1 x 10 stock, you can vary the height, width and shelf spacing to suit your needs. Keep in mind, though, that extending the width of the cabinet may require the addition of central shelf supports.",
// 	"estimatedTime": "12 hours",
// 	"materialsNeeded": "* 1/2 x 3/4 inch parting strip\n* 1 x 2 common pine\n* 1 x 4 common pine\n* 1 x 10 common pine\n* 1/4 inch thick lauan plywood\n* Finishing Nails\n* Sandpaper\n* Wood Glue\n* Wood Filler\n* Minwax Oil Based Polyurethane\n",
// 	"steps": [
// 		{
// 			"_id": "57029ed4795118be119cc43f",
// 			"stepNumber": 1,
// 			"title": "Cutting the Parts",
// 			"description": "For precise crosscuts, first make a simple, self-aligning T-guide for your circular saw. Cut a piece of 1/2-in. plywood to 2 1/2 x 24 in. and glue and screw it to a roughly 12-in.-long piece of 1 x 4 pine that will serve as the crossbar of the T. Center the plywood strip along the 1 x 4 and make sure the pieces are perfectly square to each other.\n\nButt the crossbar of the T-guide against the edge of a piece of scrap lumber, tack the guide in place and make a cut through the 1 x 4 with your saw base guided by the plywood strip. Then, trim the 1 x 4 on the opposite side in the same way. Now, the ends of the 1 x 4 can be aligned with layout lines on the stock for precise cut positioning.\n\nBegin construction by using a tape measure to mark the length of a side panel on 1 x 10 stock, and lay out the cut line with a square. The side panels on our bookcase are 48 in. long."
// 		},
// 		{
// 			"_id": "57029ed4795118be119cc43e",
// 			"stepNumber": 2,
// 			"title": "Blah Blah Blah",
// 			"description": "And some other stuff..."
// 		}
// 	],
// 	"reviews": [
// 		"->reviews.one",
// 		"->reviews.two"
// 	]
// },
// "two": {
// 	"_id": "57029ed4795118be119cc440",
// 	"user": "->users.joesmith",
// 	"title": "Learn How to Program",
// 	"description": "In this course, you'll learn how to write code like a pro!",
// 	"estimatedTime": "6 hours",
// 	"materialsNeeded": "* Notebook computer running Mac OS X or Windows\n* Text editor",
// 	"steps": [
// 		{
// 			"_id": "57029ed4795118be119cc442",
// 			"stepNumber": 1,
// 			"title": "Blah Blah Blah",
// 			"description": "And some stuff..."
// 		},
// 		{
// 			"_id": "57029ed4795118be119cc441",
// 			"stepNumber": 2,
// 			"title": "Blah Blah Blah",
// 			"description": "And some other stuff..."
// 		}
// 	],
// 	"reviews": [
// 		"->reviews.three"
// 	]
// }
