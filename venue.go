package foursquarego

type Venue struct {
	ID            string     `json:"id"`
	Name          string     `json:"name"`
	Contact       Contact    `json:"contact"`
	Location      Location   `json:"location"`
	CanonicalUrl  string     `json:"canonicalUrl"`
	Categories    []Category `json:"categories"`
	Verified      bool       `json:"verified"`
	Stats         Stats      `json:"stats"`
	Url           string     `json:"url"`
	Price         Price      `json:"price"`
	HasMenu       bool       `json:"hasMenu"`
	Likes         Likes      `json:"likes"`
	Like          bool       `json:"like"`
	Dislike       bool       `json:"dislike"`
	Ok            bool       `json:"ok"`
	Rating        float32    `json:"rating"`
	RatingColor   string     `json:"ratingColor"`
	RatingSignals int        `json:"ratingSignals"`
	Menu          ShortMenu  `json:"menu"`
	Specials      Specials   `json:"specials"`
	Photos        Photos     `json:"photos"`
	Reasons       Reasons    `json:"reasons"`
	CreatedAt     int        `json:"createdAt"`
	Tips          Tips       `json:"tips"`
	Tags          []string   `json:"tags"`
	ShortUrl      string     `json:"shortUrl"`
	TimeZone      string     `json:"timeZone"`
	Listed        Listed     `json:"listed"`
	Phrases       []Phrase   `json:"phrases"`
	Hours         HoursVenue `json:"hours"`
	Popular       Popular    `json:"popular"`
	PageUpdates   Omit       `json:"-"`
	Inbox         Omit       `json:"-"`
	VenueChains   Omit       `json:"-"`
	Attributes    Attributes `json:"attributes"`
	BestPhoto     Photo      `json:"bestPhoto"`
}

func (v *Venue) HasCategory(cat string) bool {
	for _, c := range v.Categories {
		if c.Name == cat {
			return true
		}
	}
	return false
}

type Contact struct {
	Phone          string `json:"phone"`
	FormattedPhone string `json:"formattedPhone"`
	Twitter        string `json:"twitter"`
}

type Location struct {
	Address          string   `json:"address"`
	CrossStreet      string   `json:"crossStreet"`
	Lat              float64  `json:"lat"`
	Lng              float64  `json:"lng"`
	Distance         int      `json:"distance"`
	PostalCode       string   `json:"postalCode"`
	Cc               string   `json:"cc"`
	City             string   `json:"city"`
	State            string   `json:"state"`
	Country          string   `json:"country"`
	FormattedAddress []string `json:"formattedAddress"`
	IsFuzzed         bool     `json:"isFuzzed"`
}

type Category struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	PluralName string `json:"pluralName"`
	ShortName  string `json:"shortName"`
	Icon       Icon   `json:"icon"`
	Primary    bool   `json:"primary"`
}

type Icon struct {
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
}

type Stats struct {
	CheckinsCount int `json:"checkinsCount"`
	UsersCount    int `json:"usersCount"`
	TipCount      int `json:"tipCount"`
	VisitsCount   int `json:"visitsCount"`
}

type Price struct {
	Tier     int    `json:"tier"`
	Message  string `json:"message"`
	Currency string `json:"currency"`
}

type Likes struct {
	Count   int    `json:"count"`
	Groups  Omit   `json:"-"` //TODO: take care fo this later
	Items   []User `json:"items"`
	Summary string `json:"summary"`
}

type ShortMenu struct {
	Type      string `json:"type"`
	Label     string `json:"lablel"`
	Anchor    string `json:"anchor"`
	Url       string `json:"url"`
	MobileUrl string `json:"mobileUrl"`
}

type Menu struct {
	MenuID  string  `json:"menuId"`
	Name    string  `json:"name"`
	Entries Entries `json:"entries"`
}

type Entries struct {
	Count int         `json:"count"`
	Items []EntryItem `json:"items"`
}

type EntryItem struct {
	EntryID     string   `json:"entryId"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Prices      []string `json:"prices"`
	Price       string   `json:"price"`
	Options     Omit     `json:"-"`
	Additions   Omit     `json:"-"`
}

type Specials struct {
	Count int           `json:"count"`
	Items []SpecialItem `json:"items"`
}

type SpecialItem struct {
	id                  string      `json:"id"`
	Type                string      `json:"type"`
	Message             string      `json:"message"`
	Description         string      `json:"description"`
	FinePrint           string      `json:"finePrint"`
	Unlocked            bool        `json:"unlocked"`
	Icon                string      `json:"icon"`
	Title               string      `json:"title"`
	State               string      `json:"state"`
	Provider            string      `json:"provider"`
	Redemption          string      `json:"redemption"`
	Interaction         Interaction `json:"interaction"`
	Progress            int         `json:"progress"`
	ProgressDescription string      `json:"progressDescription"`
	Detail              int         `json:"detail"`
	Target              int         `json:"target"`
	FriendsHere         []User      `json:"friendsHere"`
}

type Interaction struct {
	EntryUrl string `json:"entryUrl"`
}

type Photos struct {
	Count  int          `json:"count"`
	Groups []PhotoGroup `json:"groups"`
}

type PhotoGroup struct {
	Type  string  `json:"type"`
	Name  string  `json:"name"`
	Count int     `json:"count"`
	Items []Photo `json:"items"`
}

type Photo struct {
	ID         string      `json:"id"`
	CreatedAt  int         `json:"createdAt"`
	Source     PhotoSource `json:"source"`
	Prefix     string      `json:"prefix"`
	Suffix     string      `json:"suffix"`
	Width      int         `json:"width"`
	Height     int         `json:"height"`
	User       User        `json:"user"`
	Visibility string      `json:"visibility"`
}

type PhotoSource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type User struct {
	ID           string    `json:"id"`
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	Gender       string    `json:"gender"`
	RelationShip string    `json:"relationship"`
	Photo        UserPhoto `json:"photo"`
}

type UserPhoto struct {
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
}

type HereNow struct {
	Count   int           `json:"count"`
	Summary string        `json:"summary"`
	Groups  Omit          `json:"groups,omitempty"` //TODO: take care fo this later
	Items   []HereNowItem `json:"items,omitempty"`  // I dunno sometimes it has groups or Items
}

type HereNowItem struct {
	ID             string `json:"id"`
	CreatedAt      int    `json:"createdAt"`
	Type           string `json:"type"`
	TimeZoneOffset int    `json:"timeZoneOffset"`
	User           User   `json:"user"`
	Likes          Likes  `json:"likes"`
	Like           bool   `json:"like"`
}

type Reasons struct {
	Count int      `json:"count"`
	Items []Reason `json:"items"`
}

type Reason struct {
	Summary    string `json:"summary"`
	Type       string `json:"type"`
	ReasonName string `json:"reasonName"`
}

type Tips struct {
	Count  int        `json:"count"`
	Groups []TipGroup `json:"groups"`
}

type TipGroup struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Count int    `json:"count"`
	Items []Tip  `json:"items"`
}

type Tip struct {
	ID           string `json:"id"`
	CreatedAt    int    `json:"createdAt"`
	Text         string `json:"text"`
	Type         string `json:"type"`
	Url          string `json:"url"`
	CanonicalUrl string `json:"canonicalurl"`
	Photo        Photo  `json:"photo"`
	PhotoUrl     string `json:"photoUrl"`
	Flags        Omit   `json:"flags"` //TODO:  take care fo this later
	Likes        Likes  `json:"likes"`
	Like         bool   `json:"like"`
	LogView      bool   `json:"logView"`
	EditedAt     int    `json:"editedAt"`
	Todo         Omit   `json:"todo"` //TODO:  take care fo this later
	User         User   `json:"user"`
}

type Listed struct {
	Count  int         `json:"count"`
	Groups []ListGroup `json:"groups"`
}

type ListGroup struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Count int    `json:"count"`
	Items []List `json:"items"`
}

type List struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Type          string    `json:"type"`
	User          User      `json:"user"`
	Editable      bool      `json:"editable"`
	Public        bool      `json:"public"`
	Collaborative bool      `json:"collaborative"`
	Url           string    `json:"url"`
	CanonicalUrl  string    `json:"canonicalUrl"`
	CreatedAt     int       `json:"createdAt"`
	UpdatedAt     int       `json:"updatedAt"`
	Photo         Photo     `json:"photo"`
	Followers     Followers `json:"followers"`
	ListItems     ListItems `json:"listItems"`
}

type Followers struct {
	Count int `json:"count"`
}

type ListItems struct {
	Count int        `json:"count"`
	Items []ListItem `json:"items"`
}

type ListItem struct {
	ID        string `json:"id"`
	CreatedAt int    `json:"createdAt"`
	Tip       Tip    `json:"tip"`
	Photo     Photo  `json:"photo"`
}

type Phrase struct {
	Phrase string `json:"phrase"`
	Sample Sample `json:"sample"`
	Count  int    `json:"count"`
}

type Sample struct {
	Entities []Entity `json:"entities"`
	Text     string   `json:"text"`
	Count    int      `json:"count"`
}

type Entity struct {
	Indices []int  `json:"indices"`
	Type    string `json:"type"`
}

type HoursVenue struct {
	Status     string      `json:"status,omitempty"`
	IsOpen     bool        `json:"isOpen,ommitempty"`
	Timeframes []TimeFrame `json:"timeframes,omitempty"`
}

type Hours struct {
	Status     string           `json:"status,omitempty"`
	IsOpen     bool             `json:"isOpen,ommitempty"`
	Timeframes []TimeFrameHours `json:"timeframes,omitempty"`
}

// This is used in the main Hours struct for the entire venu
type TimeFrame struct {
	Days          string `json:"days"`
	IncludesToday bool   `json:"includesToday"`
	Open          []Open `json:"open"`
	Segments      Omit   `json:"-"` //TODO:  take care fo this later
}

// This is used for the hours endpoint
type TimeFrameHours struct {
	Days          []int  `json:"days"`
	IncludesToday bool   `json:"includesToday"`
	Open          []Open `json:"open"`
	Segments      Omit   `json:"-"`
}

type Open struct {
	Start        string `json:"start"`
	End          string `json:"end"`
	RenderedTime string `json:"renderedTime"`
}

type Popular struct {
	Status     string      `json:"status"`
	IsOpen     bool        `json:"isOpen"`
	TimeFrames []TimeFrame `json:"timeFrames"`
}

type Attributes struct {
	Groups []Attribute `json:"groups"`
}

type Attribute struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Summary string `json:"summary"`
	Count   int    `json:"count"`
	Items   []Omit `json:"-"`
}

type Event struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Categories []Category `json:"categories"`
	HereNow    HereNow    `json:"hereNow"`
	AllDay     bool       `json:"allDay"`
	Date       int        `json:"date"`
	TimeZone   string     `json:"timeZone"`
	Stats      Stats      `json:"stats"`
	Url        string     `json:"url"`
}

type Link struct {
	Provider Provider `json:"provider"`
	LinkedID string   `json:"linkedId"`
	Url      string   `json:"url"`
}

type Provider struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	AttributionImage string `json:"attributionImage"`
	AttributionLink  string `json:"attributionLink"`
	AttributionText  string `json:"attributionText"`
}
