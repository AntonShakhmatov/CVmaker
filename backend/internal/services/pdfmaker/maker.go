package pdfmaker

import (
	bub "backend/req/http/build"
	"backend/req/http/contacts"
	"backend/req/http/experience"
	"backend/req/http/header"
	"backend/req/http/language"
	"backend/req/http/links"
	"backend/req/http/projects"
	"backend/req/http/responsibilities"
	"backend/req/http/technologies"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/go-pdf/fpdf"
	"github.com/jung-kurt/gofpdf"
)

type Maker struct {
	db *sql.DB
}

func NewMaker(db *sql.DB) *Maker {
	return &Maker{db: db}
}

func (maker *Maker) LoadLanguage(ctx context.Context, id int64) (*language.LanguageForm, error) {
	query := `
		SELECT 
			language
		FROM resumes
		WHERE id = ?
	`
	row := maker.db.QueryRowContext(ctx, query, id)

	var res language.LanguageForm

	err := row.Scan(
		&res.Language,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (maker *Maker) LoadHeader(ctx context.Context, id int64) (*header.HeaderForm, error) {
	query := `
		SELECT
			first_name,
			last_name,
			competence,
			title
		FROM headers
		WHERE id = ?
	`

	row := maker.db.QueryRowContext(ctx, query, id)

	var res header.HeaderForm

	err := row.Scan(
		&res.FirstName,
		&res.LastName,
		&res.Competence,
		&res.Title,
	)

	if err != nil {
		return nil, err
	}

	return &res, err
}

func (maker *Maker) LoadLastname(ctx context.Context, id int64) (*header.HeaderForm, error) {
	query := `
		SELECT
			last_name
		FROM headers
		WHERE id = ?
	`
	row := maker.db.QueryRowContext(ctx, query, id)

	var res header.HeaderForm

	err := row.Scan(
		&res.LastName,
	)

	return &res, err
}

func (m *Maker) MakeCvHeader(outputPath string, data *header.HeaderForm) error {
	pdf := fpdf.New(fpdf.OrientationPortrait, fpdf.UnitPoint, fpdf.PageSizeA4, "")
	pdf.AddPage()
	pdf.SetFont("Courier", "", 12)

	pdf.Cell(0, 10, "My Info:")
	pdf.SetFont("Courier", "", 12)
	pdf.SetFontSize(22)
	pdf.Ln(12)

	pdf.SetFont("Courier", "", 12)
	pdf.SetFontSize(22)
	pdf.Ln(12)

	pdf.Cell(60, 8, "Name:")
	pdf.Cell(0, 8, data.FirstName)
	pdf.Ln(22)

	pdf.Cell(60, 8, "Lastname:")
	pdf.Cell(0, 8, data.LastName)
	pdf.Ln(22)

	pdf.Cell(60, 8, "Competence:")
	pdf.Cell(0, 8, data.Competence)
	pdf.Ln(22)

	pdf.Cell(60, 8, "Title:")
	pdf.Cell(0, 8, data.Title)
	pdf.Ln(22)

	return pdf.OutputFileAndClose(outputPath)
}

func (maker *Maker) LoadContact(ctx context.Context, id int64) (*contacts.ContactsForm, error) {
	query := `
		SELECT
			phone,
			email,
			address_field
		FROM contacts
		WHERE id = ?
	`

	row := maker.db.QueryRowContext(ctx, query, id)

	var res contacts.ContactsForm

	err := row.Scan(
		&res.Phone,
		&res.Email,
		&res.Address,
	)

	if err != nil {
		return nil, err
	}

	return &res, err
}

func (m *Maker) MakeCvContacts(outputPath string, data *contacts.ContactsForm) error {
	pdf := fpdf.New(fpdf.OrientationPortrait, fpdf.UnitPoint, fpdf.PageSizeA4, "")
	pdf.AddPage()
	pdf.SetFont("Courier", "", 12)

	pdf.Cell(0, 10, "My Info:")
	pdf.SetFont("Courier", "", 12)
	pdf.SetFontSize(22)
	pdf.Ln(12)

	pdf.SetFont("Courier", "", 12)
	pdf.SetFontSize(22)
	pdf.Ln(12)

	pdf.Cell(60, 8, "Phone:")
	pdf.Cell(0, 8, data.Phone)
	pdf.Ln(22)

	pdf.Cell(60, 8, "Email:")
	pdf.Cell(0, 8, data.Email)
	pdf.Ln(22)

	pdf.Cell(60, 8, "Address:")
	pdf.Cell(0, 8, data.Address)
	pdf.Ln(22)

	return pdf.OutputFileAndClose(outputPath)
}

func (maker *Maker) LoadExperience(ctx context.Context, id int64) (*experience.ExperienceForm, error) {
	query := `
		SELECT
			company,
			location_field,
			position,
			date_from,
			date_to
		FROM experience
		WHERE id = ?
	`

	row := maker.db.QueryRowContext(ctx, query, id)

	var res experience.ExperienceForm

	err := row.Scan(
		&res.Company,
		&res.Location_field,
		&res.Position,
		&res.Date_from,
		&res.Date_to,
	)

	if err != nil {
		return nil, err
	}

	return &res, err
}

func (m *Maker) MakeCvExperience(outputPath string, data *experience.ExperienceForm) error {
	pdf := fpdf.New(fpdf.OrientationPortrait, fpdf.UnitPoint, fpdf.PageSizeA4, "")
	pdf.AddPage()
	pdf.SetFont("Courier", "", 12)

	pdf.Cell(0, 10, "My Info:")
	pdf.SetFont("Courier", "", 12)
	pdf.SetFontSize(22)
	pdf.Ln(12)

	pdf.SetFont("Courier", "", 12)
	pdf.SetFontSize(22)
	pdf.Ln(12)

	pdf.Cell(60, 8, "Company:")
	pdf.Cell(0, 8, data.Company)
	pdf.Ln(22)

	pdf.Cell(60, 8, "Location field:")
	pdf.Cell(0, 8, data.Location_field)
	pdf.Ln(22)

	pdf.Cell(60, 8, "Position:")
	pdf.Cell(0, 8, data.Position)
	pdf.Ln(22)

	pdf.Cell(60, 8, "Date from:")
	pdf.Cell(0, 8, data.Date_from)
	pdf.Ln(22)

	pdf.Cell(60, 8, "Date to:")
	pdf.Cell(0, 8, data.Date_to)
	pdf.Ln(22)

	return pdf.OutputFileAndClose(outputPath)
}

func (maker *Maker) LoadProject(ctx context.Context, id int64) (*projects.ProjectsForm, error) {
	query := `
		SELECT
			name,
			description
		FROM projects
		WHERE id = ?
	`

	row := maker.db.QueryRowContext(ctx, query, id)

	var res projects.ProjectsForm

	err := row.Scan(
		&res.Name,
		&res.Description,
	)

	if err != nil {
		return nil, err
	}

	return &res, err
}

func (m *Maker) MakeCvProjects(outputPath string, data *projects.ProjectsForm) error {
	pdf := fpdf.New(fpdf.OrientationPortrait, fpdf.UnitPoint, fpdf.PageSizeA4, "")
	pdf.AddPage()
	pdf.SetFont("Courier", "", 12)

	pdf.Cell(0, 10, "My Info:")
	pdf.SetFont("Courier", "", 12)
	pdf.SetFontSize(22)
	pdf.Ln(12)

	pdf.SetFont("Courier", "", 12)
	pdf.SetFontSize(22)
	pdf.Ln(12)

	pdf.Cell(60, 8, "Name:")
	pdf.Cell(0, 8, data.Name)
	pdf.Ln(22)

	pdf.Cell(60, 8, "Location field:")
	pdf.Cell(0, 8, data.Description)
	pdf.Ln(22)

	return pdf.OutputFileAndClose(outputPath)
}

func (maker *Maker) LoadResponsibilities(ctx context.Context, id int64) (*responsibilities.ResponsibilitiesForm, error) {
	query := `
		SELECT
			responsibilities_description
		FROM responsibilities
		WHERE id = ?
	`

	row := maker.db.QueryRowContext(ctx, query, id)

	var res responsibilities.ResponsibilitiesForm

	err := row.Scan(
		&res.Description,
	)

	if err != nil {
		return nil, err
	}

	return &res, err
}

func (m *Maker) MakeCvResponsibilities(outputPath string, data *responsibilities.ResponsibilitiesForm) error {
	pdf := fpdf.New(fpdf.OrientationPortrait, fpdf.UnitPoint, fpdf.PageSizeA4, "")
	pdf.AddPage()
	pdf.SetFont("Courier", "", 12)

	pdf.Cell(0, 10, "My Info:")
	pdf.SetFont("Courier", "", 12)
	pdf.SetFontSize(22)
	pdf.Ln(12)

	pdf.SetFont("Courier", "", 12)
	pdf.SetFontSize(22)
	pdf.Ln(12)

	pdf.Cell(60, 8, "Responsibilities:")
	pdf.Cell(0, 8, data.Description)
	pdf.Ln(22)

	return pdf.OutputFileAndClose(outputPath)
}

func (maker *Maker) LoadTechnologies(ctx context.Context, id int64) (*technologies.TechnologiesForm, error) {
	query := `
		SELECT
			technologies_name
		FROM technologies
		WHERE id = ?
	`

	row := maker.db.QueryRowContext(ctx, query, id)

	var res technologies.TechnologiesForm

	err := row.Scan(
		&res.Name,
	)

	if err != nil {
		return nil, err
	}

	return &res, err
}

func (m *Maker) MakeCvTechnologies(outputPath string, data *technologies.TechnologiesForm) error {
	pdf := fpdf.New(fpdf.OrientationPortrait, fpdf.UnitPoint, fpdf.PageSizeA4, "")
	pdf.AddPage()
	pdf.SetFont("Courier", "", 12)

	pdf.Cell(0, 10, "My Info:")
	pdf.SetFont("Courier", "", 12)
	pdf.SetFontSize(22)
	pdf.Ln(12)

	pdf.SetFont("Courier", "", 12)
	pdf.SetFontSize(22)
	pdf.Ln(12)

	pdf.Cell(60, 8, "Name of technologies:")
	pdf.Cell(0, 8, data.Name)
	pdf.Ln(22)

	return pdf.OutputFileAndClose(outputPath)
}

func (maker *Maker) LoadLinks(ctx context.Context, id int64) (*links.LinksForm, error) {
	query := `
		SELECT
			label
		FROM links
		WHERE id = ?
	`

	row := maker.db.QueryRowContext(ctx, query, id)

	var res links.LinksForm

	err := row.Scan(
		&res.Label,
	)

	if err != nil {
		return nil, err
	}

	return &res, err
}

func (m *Maker) MakeCvLinks(outputPath string, data *links.LinksForm) error {
	pdf := fpdf.New(fpdf.OrientationPortrait, fpdf.UnitPoint, fpdf.PageSizeA4, "")
	pdf.AddPage()
	pdf.SetFont("Courier", "", 12)

	pdf.Cell(0, 10, "My Info:")
	pdf.SetFont("Courier", "", 12)
	pdf.SetFontSize(22)
	pdf.Ln(12)

	pdf.SetFont("Courier", "", 12)
	pdf.SetFontSize(22)
	pdf.Ln(12)

	pdf.Cell(60, 8, "Github link:")
	pdf.Cell(0, 8, data.Label)
	pdf.Ln(22)

	return pdf.OutputFileAndClose(outputPath)
}

func (m *Maker) LoadAll(ctx context.Context, cvID int64) (*bub.BuildForm, error) {
	header, err := m.LoadHeader(ctx, cvID)
	if err != nil {
		return nil, err
	}

	contacts, err := m.LoadContact(ctx, cvID)
	if err != nil {
		return nil, err
	}

	experience, err := m.LoadExperience(ctx, cvID)
	if err != nil {
		return nil, err
	}

	links, err := m.LoadLinks(ctx, cvID)
	if err != nil {
		return nil, err
	}

	project, err := m.LoadProject(ctx, cvID)
	if err != nil {
		return nil, err
	}

	responsibilities, err := m.LoadResponsibilities(ctx, cvID)
	if err != nil {
		return nil, err
	}

	technologies, err := m.LoadTechnologies(ctx, cvID)
	if err != nil {
		return nil, err
	}

	return &bub.BuildForm{
		Header:           header,
		Contacts:         contacts,
		Experience:       experience,
		Links:            links,
		Projects:         project,
		Responsibilities: responsibilities,
		Technologies:     technologies,
	}, nil
}

func (m *Maker) MakeCV(outPath string, d *bub.BuildForm) error {
	pdf := gofpdf.New(gofpdf.OrientationPortrait, gofpdf.UnitPoint, gofpdf.PageSizeA4, "")
	pdf.SetTitle("CV", false)
	pdf.SetAuthor("CV MAKER", false)
	pdf.SetMargins(48, 56, 48)
	pdf.SetAutoPageBreak(true, 56)

	fontBody := "Helvetica"
	fontMono := "Courier"

	pdf.SetHeaderFunc(func() {
		pdf.SetFont(fontBody, "B", 12)
		pdf.SetTextColor(20, 20, 20)

		pdf.CellFormat(0, 16, "CV", "", 0, "L", false, 0, "")

		pdf.SetFont(fontBody, "", 9)
		pdf.SetTextColor(90, 90, 90)
		pdf.CellFormat(0, 16, time.Now().Format("2006-01-02"), "", 0, "R", false, 0, "")

		pdf.Ln(18)
		x1, y := pdf.GetX(), pdf.GetY()
		pdf.SetDrawColor(210, 210, 210)
		pdf.Line(x1, y, 595-48, y)
		pdf.Ln(14)

		pdf.SetTextColor(20, 20, 20)
	})

	pdf.SetFooterFunc(func() {
		pdf.SetY(-44)
		pdf.SetDrawColor(210, 210, 210)
		x1 := 48.0
		x2 := 595.0 - 48.0
		y := pdf.GetY()
		pdf.Line(x1, y, x2, y)
		pdf.Ln(10)

		pdf.SetFont(fontBody, "", 9)
		pdf.SetTextColor(120, 120, 120)
		pdf.CellFormat(0, 14, fmt.Sprintf("Page %d/{nb}", pdf.PageNo()), "", 0, "R", false, 0, "")
	})

	pdf.AliasNbPages("")
	pdf.AddPage()

	coverPairs := []kv{
		// {"Document", "WPS"},
		// {"Generated", time.Now().Format("2006-01-02 15:04")},
	}
	drawInfoBlock(pdf, fontBody, coverPairs)

	drawSectionFromStruct(pdf, fontBody, fontMono, "My info:", d.Header)
	drawSectionFromStruct(pdf, fontBody, fontMono, "Contacts:", d.Contacts)
	drawSectionFromStruct(pdf, fontBody, fontMono, "My experience:", d.Experience)
	drawSectionFromStruct(pdf, fontBody, fontMono, "Links:", d.Links)
	drawSectionFromStruct(pdf, fontBody, fontMono, "My projects info:", d.Projects)
	drawSectionFromStruct(pdf, fontBody, fontMono, "My responsibilities in company:", d.Responsibilities)
	drawSectionFromStruct(pdf, fontBody, fontMono, "Technologies I have used:", d.Technologies)

	return pdf.OutputFileAndClose(outPath)
}

type kv struct {
	K string
	V string
}

func drawInfoBlock(pdf *gofpdf.Fpdf, font string, pairs []kv) {
	pdf.SetFillColor(245, 245, 245)
	pdf.SetDrawColor(220, 220, 220)
	pdf.SetTextColor(20, 20, 20)

	x := pdf.GetX()
	y := pdf.GetY()
	w := 595.0 - 48.0 - 48.0
	h := float64(18 + len(pairs)*14 + 10)

	pdf.Rect(x, y, w, h, "FD")

	pdf.SetFont(font, "B", 12)
	pdf.SetXY(x+12, y+12)
	pdf.CellFormat(0, 14, "My portfolio", "", 1, "L", false, 0, "")

	pdf.SetFont(font, "", 10)
	pdf.SetXY(x+12, y+30)

	for _, p := range pairs {
		pdf.SetFont(font, "B", 10)
		pdf.CellFormat(120, 14, p.K, "", 0, "L", false, 0, "")
		pdf.SetFont(font, "", 10)
		pdf.CellFormat(0, 14, safeVal(p.V), "", 1, "L", false, 0, "")
	}

	pdf.Ln(18)
}

func drawSectionFromStruct(pdf *gofpdf.Fpdf, fontBody, fontMono, title string, data any) {
	if isNilInterface(data) {
		return
	}

	drawSectionTitle(pdf, fontBody, title)

	pairs := structToPairs(data)

	drawKeyValueTable(pdf, fontBody, fontMono, pairs)

	pdf.Ln(14)
}

func drawSectionTitle(pdf *gofpdf.Fpdf, font string, title string) {
	pdf.SetFillColor(35, 35, 35)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetFont(font, "B", 11)
	pdf.CellFormat(0, 18, "  "+title, "", 1, "L", true, 0, "")
	pdf.Ln(8)

	pdf.SetTextColor(20, 20, 20)
}

func drawKeyValueTable(pdf *gofpdf.Fpdf, fontBody, fontMono string, pairs []kv) {
	if len(pairs) == 0 {
		pdf.SetFont(fontBody, "", 10)
		pdf.SetTextColor(120, 120, 120)
		pdf.CellFormat(0, 14, "No data", "", 1, "L", false, 0, "")
		pdf.SetTextColor(20, 20, 20)
		return
	}

	usableW := 595.0 - 48.0 - 48.0
	labelW := 180.0
	valueW := usableW - labelW

	rowH := 16.0

	for i, p := range pairs {

		if i%2 == 0 {
			pdf.SetFillColor(250, 250, 250)
		} else {
			pdf.SetFillColor(255, 255, 255)
		}
		pdf.SetDrawColor(230, 230, 230)

		pdf.SetFont(fontMono, "", 9)
		lines := pdf.SplitLines([]byte(safeVal(p.V)), valueW-16)
		h := float64(len(lines)) * rowH
		if h < rowH {
			h = rowH
		}

		if pdf.GetY()+h > (842.0 - 56.0) {
			pdf.AddPage()
		}

		x := pdf.GetX()
		y := pdf.GetY()

		pdf.SetFont(fontBody, "B", 9)
		pdf.SetTextColor(60, 60, 60)
		pdf.Rect(x, y, labelW, h, "FD")
		pdf.SetXY(x+8, y+5)
		pdf.MultiCell(labelW-16, 12, humanizeKey(p.K), "", "L", false)

		pdf.SetFont(fontMono, "", 9)
		pdf.SetTextColor(20, 20, 20)
		pdf.SetXY(x+labelW, y)
		pdf.Rect(x+labelW, y, valueW, h, "FD")
		pdf.SetXY(x+labelW+8, y+5)
		pdf.MultiCell(valueW-16, 12, safeVal(p.V), "", "L", false)

		pdf.SetXY(x, y+h)
	}
}

func structToPairs(v any) []kv {
	b, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	var m map[string]any
	if err := json.Unmarshal(b, &m); err != nil {
		return nil
	}

	orderedKeys := fieldOrder(v)

	if len(orderedKeys) == 0 {
		pairs := make([]kv, 0, len(m))
		for k, val := range m {
			pairs = append(pairs, kv{K: k, V: anyToString(val)})
		}
		return pairs
	}

	pairs := make([]kv, 0, len(orderedKeys))
	for _, k := range orderedKeys {
		if val, ok := m[k]; ok {
			pairs = append(pairs, kv{K: k, V: anyToString(val)})
		}
	}
	return pairs
}

func fieldOrder(v any) []string {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Pointer {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return nil
	}

	rt := rv.Type()
	out := make([]string, 0, rt.NumField())
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		tag := f.Tag.Get("json")
		if tag != "" {
			name := strings.Split(tag, ",")[0]
			if name != "" && name != "-" {
				out = append(out, name)
				continue
			}
		}
		out = append(out, f.Name)
	}
	return out
}

func anyToString(v any) string {
	if v == nil {
		return "—"
	}
	switch t := v.(type) {
	case string:
		if strings.TrimSpace(t) == "" {
			return "—"
		}
		return t
	case float64:
		if t == float64(int64(t)) {
			return fmt.Sprintf("%d", int64(t))
		}
		return fmt.Sprintf("%v", t)
	case bool:
		if t {
			return "Yes"
		}
		return "No"
	default:
		b, err := json.Marshal(t)
		if err != nil {
			return fmt.Sprintf("%v", t)
		}
		s := string(b)
		if s == "null" || s == `""` {
			return "—"
		}
		return s
	}
}

func safeVal(s string) string {
	if strings.TrimSpace(s) == "" {
		return "-"
	}
	return s
}

func humanizeKey(k string) string {
	k = strings.TrimSpace(k)
	if k == "" {
		return ""
	}

	if strings.Contains(k, "_") {
		parts := strings.Split(k, "_")
		for i := range parts {
			parts[i] = strings.Title(strings.ToLower(parts[i]))
		}
		return strings.Join(parts, " ")
	}

	var out []rune
	for i, r := range k {
		if i > 0 && r >= 'A' && r <= 'Z' {
			out = append(out, ' ')
		}
		out = append(out, r)
	}
	return string(out)
}

func isNilInterface(v any) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Pointer, reflect.Interface, reflect.Slice, reflect.Map:
		return rv.IsNil()
	default:
		return false
	}
}
