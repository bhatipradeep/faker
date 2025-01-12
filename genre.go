package faker

var genres = map[string]string{
	"Abimegender":    "a gender that is profound, deep, and infinite; meant to resemble when one mirror is reflecting into another mirror creating an infinite paradox.",
	"Adamasgender":   "a gender that refuses to be categorized.",
	"Aerogender":     "a gender that is influenced by your surroundings.",
	"Aesthetigender": "a gender that is derived from an aesthetic; also known as videgender.",
	"Affectugender":  "a gender that is affected by mood swings.",
	"Agender":        "the feeling of no gender/absence of gender or neutral gender.",
	"Agenderflux":    "being mostly agender except having small shifts towards other genders making them demigenders (because of the constancy of being agender).",
	"Alexigender":    "a gender that is fluid between more than one gender but the individual cannot tell what those genders are.",
	"Aliusgender":    "a gender that is removed from common gender descriptors and guidelines.",
	"Amaregender":    "a gender that changes depending on who you’re in love with.",
	"Ambigender":     "defined as having the feeling of two genders simultaneously without fluctuation; meant to reflect the concept of being ambidextrous, only with gender.",
	"Ambonec":        "identifying as both man and woman, yet neither at the same time.",
	"Amicagender":    "a gender that changes depending on which friend you’re with.",
	"Androgyne":      "sometimes used in the case of “androgynous presentation”; describes the feeling of being a mix of both masculine and feminine (and sometimes neutral) gender qualities.",
	"Anesigender":    "feeling like a certain gender yet being more comfortable identifying with another.",
	"Angenital":      "a desire to be without primary sexual characteristics, without necessarily being genderless; one may be both angenital and identify as any other gender alongside.",
	"Anogender":      "a gender that fades in and out but always comes back to the same feeling.",
	"Anongender":     "a gender that is unknown to both yourself and others.",
	"Antegender":     "a protean gender which has the potential to be anything, but is formless and motionless, and therefore, does not manifest as any particular gender.",
	"Anxiegender":    "a gender that is affected by anxiety.",
	"Apagender":      "a feeling of apathy towards one’s gender which leads to them not looking any further into it.",
	"Apconsugender":  "a gender where you know what it isn’t, but not what it is; the gender is hiding itself from you.",
	"Astergender":    "a gender that feels bright and celestial.",
	"Astralgender":   "a gender that feels connected to space.",
	"Autigender":     "a gender that can only be understood in the context of being autistic (POSSIBLE TRIGGER WARNING).",
	"Autogender":     "a gender experience that is deeply personal to oneself.",
	"Axigender":      "when a person experiences two genders that sit on opposite ends of an axis; one being agender and the other being any other gender; these genders are experienced one at a time with no overlapping and with very short transition time.",
	"Bigender":       "the feeling of having two genders either at the same time or separately; usually used to describe feeling “traditionally male” and “traditionally female”, but does not have to.",
	"Biogender":      "a gender that feels connected to nature in some way.",
	"Blurgender":     "the feeling of having more than one gender that are somehow blurred together to the point of not being able to distinguish or identify individual genders; synonymous with genderfuzz.",
	"Boyflux":        "when one feels mostly or all male most of the time but experience fluctuating intensity of male identity.",
	"Burstgender":    "and gender that comes in intense bursts of feeling and quickly fades back to the original state.",
	"Caelgender":     "a gender that shares qualities with outer space or has the aesthetic of space, stars, nebulas, etc.",
	"Cassgender":     "the feeling of gender is unimportant to you.",
	"Cassflux":       "when the level of indifference towards your gender fluctuates.",
	"Cavusgender":    "for people with depression; when you feel one gender when not depressed and another when depressed.",
	"Cendgender":     "when your gender changes between one and its opposite.",
	"Ceterofluid":    "when you are ceterogender and your feelings fluctuate between masculine, feminine, and neutral.",
	"Ceterogender":   "a nonbinary gender with specific masculine, feminine, or neutral feelings.",
	"Cisgender":      "the feeling of being the gender you were assigned at birth, all the time (assigned (fe)male/feeling (fe)male).",
	"Cloudgender":    "a gender that cannot be fully realized or seen clearly due to depersonalization/derealization disorder.",
	"Collgender":     "the feeling of having too many genders simultaneously to describe each one.",
	"Colorgender":    "a gender associated with one or more colors and the feelings, hues, emotions, and/or objects associated with that color; may be used like pinkgender, bluegender, yellowgender.",
	"Commogender":    "when you know you aren’t cisgender, but you settled with your assigned gender for the time being.",
	"Condigender":    "a gender that is only felt during certain circumstances.",
	"Deliciagender":  "from the Latin word Delicia meaning “favorite”, meaning the feeling of having more than one simultaneous gender yet preferring one that fits better.",
	"Demifluid":      "the feeling your gender being fluid throughout all the demigenders; the feeling of having multiple genders, some static and some fluid.",
	"Demiflux":       "the feeling of having multiple genders, some static and some fluctuating.",
	"Demigender":     "a gender that is partially one gender and partially another.",
	"Domgender":      "having more than one gender yet one being more dominant than the others.",
	"Duragender":     "from the Latin word dura meaning “long-lasting”, meaning a subcategory of multigender in which one gender is more identifiable, long-lasting, and prominent than the other genders.",
	"Egogender":      "a gender that is so personal to your experience that it can only be described as “you”.",
	"Epicene":        "sometimes used synonymously with the adjective “androgynous”; the feeling of either having or not displaying characteristics of both or either binary gender; sometimes used to describe feminine male-identifying individuals.",
	"Espigender":     "a gender that is related to being a spirit or exists on a higher or extradimensional plane.",
	"Exgender":       "the outright refusal to accept or identify in, on, or around the gender spectrum.",
	"Existigender":   "a gender that only exists or feels present when thought about or when a conscious effort is made to notice it.",
	"Femfluid":       "having fluctuating or fluid gender feelings that are limited to feminine genders.",
	"Femgender":      "a nonbinary gender which is feminine in nature.",
	"Fluidflux":      "the feeling of being fluid between two or more genders that also fluctuate in intensity; a combination of genderfluid and genderflux.",
	"Gemigender":     "having two opposite genders that work together, being fluid and flux together.",
	"Genderblank":    "a gender that can only be described as a blank space; when gender is called into question, all that comes to mind is a blank space.",
	"Genderflow":     "a gender that is fluid between infinite feelings.",
	"Genderfluid":    "the feeling of fluidity within your gender identity; feeling a different gender as time passes or as situations change; not restricted to any number of genders.",
	"Genderflux":     "the feeling of your gender fluctuating in intensity; like genderfluid but between one gender and agender.",
	"Genderfuzz":     "coined by lolzmelmel; the feeling of having more than one gender that are somehow blurred together to the point of not being able to distinguish or identify individual genders; synonymous with blurgender.",
	"Gender Neutral": "the feeling of having a neutral gender, whether somewhere in between masculine and feminine or a third gender that is separate from the binary; often paired with neutrois.",
	"Genderpunk":     "a gender identity that actively resists gender norms.",
	"Genderqueer":    "originally used as an umbrella term for nonbinary individuals; it may be used as an identity; it describes a nonbinary gender regardless of whether the individual is masculine or feminine leaning.",
	"Genderwitched":  "a gender in which one is intrigued or entranced by the idea of a particular gender, but is not certain that they are actually feeling it.",
	"Girlflux":       "when one feels mostly or all female most of the time but experiences fluctuating intensities of female identity.",
	"Glassgender":    "a gender that is very sensitive and fragile.",
	"Glimragender":   "a faintly shining, wavering gender.",
	"Greygender":     "having a gender that is mostly outside of the binary but is weak and can barely be felt.",
	"Gyragender":     "having multiple genders but understanding none of them.",
	"Healgender":     "a gender that once realized, brings lots of peace, clarity, security, and creativity to the individual’s mind.",
	"Heliogender":    "a gender that is warm and burning.",
	"Hemigender":     "a gender that is half one gender and half something else; one or both halves may be identifiable genders.",
	"Horogender":     "a gender that changes over time with the core feeling of remaining the same.",
	"Hydrogender":    "a gender that shares qualities with water.",
	"Imperigender":   "a fluid gender that can be controlled by the individual.",
	"Intergender":    "the feeling of gender falling somewhere on the spectrum between masculine and feminine.",
	"Juxera":         "a feminine gender similar to girl, but on a separate plane and off to itself.",
	"Libragender":    "a gender that feels agender but has a strong connection to another gender.",
	"Magigender":     "a gender that is mostly gender and the rest is something else.",
	"Mascfluid":      "A gender that is fluid in nature, and restricted only to masculine genders.",
	"Mascgender":     "a non-binary gender which is masculine in nature.",
	"Maverique":      "taken from the word maverick; the feeling of having a gender that is separate from masculinity, femininity, and neutrality, but is not agender; a form of a third gender.",
	"Mirrorgender":   "a gender that changes to fit the people around you.",
	"Molligender":    "a gender that is soft, subtle, and subdued.",
	"Multigender":    "the feeling of having more than one simultaneous or fluctuating gender; simultaneous with multigender and omnigender.",
	"Nanogender":     "feeling a small part of one gender with the rest being something else.",
	"Neutrois":       "the feeling of having a neutral gender; sometimes a lack of gender that leads to feeling neutral.",
	"Nonbinary":      "originally an umbrella term for any gender outside the binary of cisgenders; may be used as an individual identity; occasionally used alongside of genderqueer.",
	"Omnigender":     "the feeling of having more than one simultaneous or fluctuating gender; simultaneous with multigender and polygender.",
	"Oneirogender":   "coined by anonymous, “being agender, but having recurring fantasies or daydreams of being a certain gender without the dysphoria or desire to actually be that gender day-to-day”.",
	"Pangender":      "the feeling of having every gender; this is considered problematic by some communities and thus has been used as the concept of relating in some way to all genders as opposed to containing every gender identity; only applies to genders within one’s own culture.",
	"Paragender":     "the feeling very near one gender and partially something else which keeps you from feeling fully that gender.",
	"Perigender":     "identifying with gender but not as a gender.",
	"Polygender":     "the feeling of having more than one simultaneous or fluctuating gender; simultaneous with multigender and omnigender.",
	"Proxvir":        "a masculine gender similar to a boy, but on a separate plane and off to itself.",
	"Quoigender":     "feeling as if the concept of gender is inapplicable or nonsensical to one’s self.",
	"Subgender":      "mostly a gender with a bit of another gender.",
	"Surgender":      "having a gender that is 100% one gender but with more of another gender added on top of that.",
	"Systemgender":   "a gender that is the sum of all the genders within a multiple or median system.",
	"Tragender":      "a gender that stretches over the whole spectrum of genders.",
	"Transgender":    "any gender identity that transcends or does not align with your assigned gender or society’s idea of gender; the feeling of being any gender that does not match your assigned gender.",
	"Trigender":      "the feeling of having three simultaneous or fluctuating genders.",
	"Vapogender":     "a gender that sort of feels like smoke; it can be seen on a shallow level but once you go deeper, it disappears and you are left with no gender and only tiny wisps of what you thought it was.",
	"Venngender":     "when two genders overlap creating an entirely new gender; like a Venn diagram.",
	"Verangender":    "a gender that seems to shift/change the moment it is identified.",
	"Vibragender":    "a gender that is usually one stable gender but will occasionally change or fluctuate before stabilizing again.",
	"Vocigender":     "a gender that is weak or hollow.",
}

// Genre is a faker struct for Genre
type Genre struct {
	Faker *Faker
}

// Name returns a genre name for Genre
func (f Genre) Name() string {
	return f.Faker.RandomStringMapKey(genres)
}

// NameWithDescription returns a name and description for Genre
func (f Genre) NameWithDescription() (string, string) {
    key := f.Faker.RandomStringMapKey(genres)
    return key, genres[key]
}
