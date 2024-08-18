package generator_test

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// ===============================================================
// Utils
// ===============================================================

func createTable(db *sql.DB, ctx context.Context, query string) error {
	_, err := db.ExecContext(ctx, query)
	return err
}

func NewEntity(db *sql.DB) *Repository {
	return &Repository{db: db}
}

type Repository struct {
	db *sql.DB
}

func (this *Repository) Label(ctx context.Context) *LabelRepository {
	return &LabelRepository{db: this.db, ctx: ctx}
}

func (this *Repository) DatedLabel(ctx context.Context) *DateLabelRepository {
	return &DateLabelRepository{db: this.db, ctx: ctx}
}

// ===============================================================
// Label
// ===============================================================

type Label struct {
	Id   uuid.UUID
	Name string
}

type LabelRepository struct {
	db  *sql.DB
	ctx context.Context
}

func (this *LabelRepository) EnsureTable(ctx context.Context) error {
	query := `CREATE TABLE IF NOT EXISTS label (id TEXT, name TEXT);`
	return createTable(this.db, ctx, query)
}

func (this *LabelRepository) Save(ctx context.Context, label Label) error {
	_, err := this.ById(ctx, label.Id)

	if err != nil {
		return nil
	}

	updateQuery := `UPDATE label SET name = ? WHERE id = ?;`

	_, err = this.db.ExecContext(ctx, updateQuery, label.Name, label.Id)

	if err != nil {
		return nil
	}

	return nil
}

func (this *LabelRepository) ById(ctx context.Context, id uuid.UUID) (Label, error) {

	return Label{}, nil
}

func (this *LabelRepository) ByName(ctx context.Context, id uuid.UUID) ([]Label, error) {

	return []Label{}, nil
}

// ===============================================================
// DateLabel
// ===============================================================

type DateLabel struct {
	Id    uuid.UUID
	Date  time.Time
	Label Label
}

type DateLabelRepository struct {
	db  *sql.DB
	ctx context.Context
}

func (this *DateLabelRepository) EnsureTable(ctx context.Context) error {
	query := `CREATE TABLE IF NOT EXISTS label (id TEXT, name TEXT);`
	return createTable(this.db, ctx, query)

	return nil
}

func (this *DateLabelRepository) Save(ctx context.Context, label DateLabel) error {

	return nil
}

func (this *DateLabelRepository) ById(ctx context.Context, id uuid.UUID) (DateLabel, error) {

	return DateLabel{}, nil
}

func (this *DateLabelRepository) ByName(ctx context.Context, id uuid.UUID) ([]DateLabel, error) {

	return []DateLabel{}, nil
}

func (this *DateLabelRepository) ByTimeRange(ctx context.Context, start, end time.Time) ([]DateLabel, error) {

	return []DateLabel{}, nil
}

// ===============================================================
// NfixType
// ===============================================================

type NfixType struct {
	Id   uuid.UUID
	Name string
}

type NfixTypeRepository struct {
	db *sql.DB
}

func (this *NfixTypeRepository) Save(ctx context.Context, label NfixType) error {

	return nil
}

func (this *NfixTypeRepository) ById(ctx context.Context, id uuid.UUID) (NfixType, error) {

	return NfixType{}, nil
}

func (this *NfixTypeRepository) ByName(ctx context.Context, id uuid.UUID) ([]NfixType, error) {

	return []NfixType{}, nil
}

func createNfixTypes() map[string]NfixType {
	common := NfixType{
		Id:   uuid.New(),
		Name: "Common",
	}

	formal := NfixType{
		Id:   uuid.New(),
		Name: "Formal",
	}

	academic := NfixType{
		Id:   uuid.New(),
		Name: "Academic",
	}

	honorary := NfixType{
		Id:   uuid.New(),
		Name: "Honorary",
	}

	professional := NfixType{
		Id:   uuid.New(),
		Name: "Professional",
	}

	social := NfixType{
		Id:   uuid.New(),
		Name: "Social",
	}

	religious := NfixType{
		Id:   uuid.New(),
		Name: "Religious",
	}

	return map[string]NfixType{
		common.Name:       common,
		formal.Name:       formal,
		academic.Name:     academic,
		honorary.Name:     honorary,
		professional.Name: professional,
		social.Name:       social,
		religious.Name:    religious,
	}
}

// ===============================================================
// Prefix
// ===============================================================

type Prefix struct {
	Id       uuid.UUID
	Name     string
	Type     NfixType
	Acronyms []string
}

type PrefixRepository struct {
	db *sql.DB
}

func (this *PrefixRepository) Save(ctx context.Context, prefix Prefix) error {

	return nil
}

func (this *PrefixRepository) ById(ctx context.Context, id uuid.UUID) (Prefix, error) {

	return Prefix{}, nil
}

func (this *PrefixRepository) ByName(ctx context.Context, id uuid.UUID) ([]Prefix, error) {

	return []Prefix{}, nil
}

func (this *PrefixRepository) ByAcronym(ctx context.Context, id uuid.UUID) ([]Prefix, error) {

	return []Prefix{}, nil
}

func (this *PrefixRepository) ByText(ctx context.Context, id uuid.UUID) ([]Prefix, error) {

	return []Prefix{}, nil
}

func createPrefixes(nfixTypes map[string]NfixType) map[uuid.UUID]Prefix {
	common := nfixTypes["Common"]
	formal := nfixTypes["Formal"]

	mr := Prefix{
		Id:       uuid.New(),
		Name:     "Mister",
		Acronyms: []string{"Mr", "Mr."},
		Type:     common,
	}

	miss := Prefix{
		Id:       uuid.New(),
		Name:     "Miss",
		Acronyms: []string{},
	}

	mrs := Prefix{
		Id:       uuid.New(),
		Name:     "Mistress",
		Acronyms: []string{"Mrs", "Mrs."},
		Type:     common,
	}

	sir := Prefix{
		Id:   uuid.New(),
		Name: "Sir",
		Type: formal,
	}

	madam := Prefix{
		Id:   uuid.New(),
		Name: "Madam",
		Type: formal,
	}

	return map[uuid.UUID]Prefix{
		mr.Id:    mr,
		miss.Id:  miss,
		mrs.Id:   mrs,
		sir.Id:   sir,
		madam.Id: madam,
	}
}

// ===============================================================
// Suffix
// ===============================================================

type Suffix struct {
	Id       uuid.UUID
	Name     string
	Type     NfixType
	Acronyms []string
}

type SuffixRepository struct {
	db *sql.DB
}

func (this *SuffixRepository) Save(ctx context.Context, prefix Suffix) error {

	return nil
}

func (this *SuffixRepository) ById(ctx context.Context, id uuid.UUID) (Suffix, error) {

	return Suffix{}, nil
}

func (this *SuffixRepository) ByName(ctx context.Context, id uuid.UUID) ([]Suffix, error) {

	return []Suffix{}, nil
}

func (this *SuffixRepository) ByAcronym(ctx context.Context, id uuid.UUID) ([]Suffix, error) {

	return []Suffix{}, nil
}

func (this *SuffixRepository) ByText(ctx context.Context, id uuid.UUID) ([]Suffix, error) {

	return []Suffix{}, nil
}

func createSuffixes(nfixTypes map[string]NfixType) map[uuid.UUID]Suffix {
	common := nfixTypes["Common"]
	academic := nfixTypes["Academic"]

	junior := Suffix{
		Id:       uuid.New(),
		Name:     "Junior",
		Acronyms: []string{"Jr", "Jr.", "Jnr"},
		Type:     common,
	}

	senior := Suffix{
		Id:       uuid.New(),
		Name:     "Senior",
		Acronyms: []string{"Sr", "Sr.", "Snr"},
	}

	ba := Suffix{
		Id:       uuid.New(),
		Name:     "Bachelor of Arts",
		Acronyms: []string{"B.A.", "A.B."},
		Type:     academic,
	}

	bed := Suffix{
		Id:       uuid.New(),
		Name:     "Bachelor of Education",
		Type:     academic,
		Acronyms: []string{"B.Ed"},
	}

	bfa := Suffix{
		Id:       uuid.New(),
		Name:     "Bachelor of Fine Arts",
		Type:     academic,
		Acronyms: []string{"B.F.A."},
	}

	bs := Suffix{
		Id:       uuid.New(),
		Name:     "Bachelor of Science",
		Type:     academic,
		Acronyms: []string{"B.S.", "B.Sc.", "B.E."},
	}

	return map[uuid.UUID]Suffix{
		junior.Id: junior,
		senior.Id: senior,
		ba.Id:     ba,
		bed.Id:    bed,
		bfa.Id:    bfa,
		bs.Id:     bs,
	}
}

// ===============================================================
// Name
// ===============================================================

type Name struct {
	Id       uuid.UUID
	Prefixes []Prefix
	Firsts   []string
	Middles  []string
	Lasts    []string
	Suffixes []Suffix
}

type NameRepository struct {
	db *sql.DB
}

func (this *NameRepository) Save(ctx context.Context, name Name) error {

	return nil
}

func (this *NameRepository) ById(ctx context.Context, id uuid.UUID) (Name, error) {

	return Name{}, nil
}

func (this *NameRepository) ByName(ctx context.Context, id uuid.UUID) ([]Name, error) {

	return []Name{}, nil
}

func (this *NameRepository) ByAcronym(ctx context.Context, id uuid.UUID) ([]Name, error) {

	return []Name{}, nil
}

func (this *NameRepository) ByText(ctx context.Context, id uuid.UUID) ([]Name, error) {

	return []Name{}, nil
}

// ===============================================================
// Person
// ===============================================================

type Person struct {
	Id     uuid.UUID
	Names  []Name
	Labels []Label
}

type PersonRepository struct {
	db *sql.DB
}

func (this *PersonRepository) Save(ctx context.Context, person Person) error {

	return nil
}

func (this *PersonRepository) ById(ctx context.Context, id uuid.UUID) (Person, error) {

	return Person{}, nil
}

func (this *PersonRepository) ByName(ctx context.Context, id uuid.UUID) ([]Person, error) {

	return []Person{}, nil
}

func (this *PersonRepository) ByAcronym(ctx context.Context, id uuid.UUID) ([]Person, error) {

	return []Person{}, nil
}

func (this *PersonRepository) ByText(ctx context.Context, id uuid.UUID) ([]Person, error) {

	return []Person{}, nil
}

func createPeople(prefixes map[string]Prefix, suffixes map[string]Suffix) map[uuid.UUID]Person {
	mr := prefixes["Mister"]
	ms := prefixes["Ms"]
	jr := suffixes["Junior"]

	john_smith := Person{
		Id: uuid.New(),
		Names: []Name{
			{
				Id:       uuid.New(),
				Prefixes: []Prefix{mr},
				Firsts:   []string{"John"},
				Lasts:    []string{"Smith"},
				Suffixes: []Suffix{jr},
			},
		},
	}

	jane_doe := Person{
		Id: uuid.New(),
		Names: []Name{
			{
				Id:       uuid.New(),
				Prefixes: []Prefix{ms},
				Firsts:   []string{"Jane"},
				Lasts:    []string{"Doe"},
			},
		},
	}

	return map[uuid.UUID]Person{
		john_smith.Id: john_smith,
		jane_doe.Id:   jane_doe,
	}
}

// ===============================================================
// PhoneNumber
// ===============================================================

type PhoneNumber struct {
	Id          uuid.UUID
	CountryCode int
	Nyx         int
	Nxx         int
	LineNumber  int
	Labels      []Label
}

type PhoneNumberRepository struct {
	db *sql.DB
}

func (this *PhoneNumberRepository) Save(ctx context.Context, person PhoneNumber) error {

	return nil
}

func (this *PhoneNumberRepository) ById(ctx context.Context, id uuid.UUID) (PhoneNumber, error) {

	return PhoneNumber{}, nil
}

func (this *PhoneNumberRepository) ByCountryCode(ctx context.Context, id uuid.UUID) ([]PhoneNumber, error) {

	return []PhoneNumber{}, nil
}

func (this *PhoneNumberRepository) ByNyx(ctx context.Context, id uuid.UUID) ([]PhoneNumber, error) {

	return []PhoneNumber{}, nil
}

func (this *PhoneNumberRepository) ByNxx(ctx context.Context, id uuid.UUID) ([]PhoneNumber, error) {

	return []PhoneNumber{}, nil
}

func (this *PhoneNumberRepository) ByLineNumber(ctx context.Context, id uuid.UUID) ([]PhoneNumber, error) {

	return []PhoneNumber{}, nil
}

func (this *PhoneNumberRepository) ByNumber(ctx context.Context, id uuid.UUID) ([]PhoneNumber, error) {

	return []PhoneNumber{}, nil
}

// ===============================================================
// Contact
// ===============================================================

type Contact struct {
	Id           uuid.UUID
	Person       Person
	PhoneNumbers []PhoneNumber
	Labels       []Label
}

type ContactRepository struct {
	db *sql.DB
}

func (this *ContactRepository) Save(ctx context.Context, person Contact) error {

	return nil
}

func (this *ContactRepository) ById(ctx context.Context, id uuid.UUID) (Contact, error) {

	return Contact{}, nil
}

func (this *ContactRepository) ByCountryCode(ctx context.Context, id uuid.UUID) ([]Contact, error) {

	return []Contact{}, nil
}

func (this *ContactRepository) ByNyx(ctx context.Context, id uuid.UUID) ([]Contact, error) {

	return []Contact{}, nil
}

func (this *ContactRepository) ByNxx(ctx context.Context, id uuid.UUID) ([]Contact, error) {

	return []Contact{}, nil
}

func (this *ContactRepository) ByLineNumber(ctx context.Context, id uuid.UUID) ([]Contact, error) {

	return []Contact{}, nil
}

func (this *ContactRepository) ByNumber(ctx context.Context, id uuid.UUID) ([]Contact, error) {

	return []Contact{}, nil
}
