package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

var (
	indexFileNmae = "000000-images.json"

	dumpBaseDir = "data"
	dumpDirName = ""

	roxAPIEndpoint = os.Getenv("ROX_ENDPOINT")
	roxAPIToken    = os.Getenv("ROX_API_TOKEN")

	roxAPIEndpointR = os.Getenv("ROX_ENDPOINTR")
	roxAPITokenR    = os.Getenv("ROX_API_TOKENR")
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	action := "dump"
	if len(os.Args) > 1 {
		action = strings.TrimSpace(os.Args[1])
	}

	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	now := time.Now()
	dumpDirName = now.Format("2006-01-02-150405")

	if action == "dump" {
		roxAPIEndpoint = strings.TrimSpace(roxAPIEndpoint)
		if roxAPIEndpoint == "" {
			panic("ROX_ENDPOINT empty")
		}

		roxAPIToken = strings.TrimSpace(roxAPIToken)
		if roxAPIToken == "" {
			panic("ROX_API_TOKEN empty")
		}

		destDir := filepath.Join(dumpBaseDir, dumpDirName)

		log.Printf("Dumping files into: %s", destDir)

		dumpAllScannerV2Scans(destDir, roxAPIEndpoint, roxAPIToken)
	}

	usage := func() {
		fmt.Printf("compare requires 3 args `go run . compare <dir1> <dir2>`\n")
		os.Exit(1)
	}

	if action == "compare" {
		if len(os.Args) < 4 {
			usage()

		}
		d1 := os.Args[2]
		d2 := os.Args[3]

		if d1 == "" || d2 == "" {
			usage()
		}

		if d1 == d2 {
			fmt.Printf("<dir1> and <dir2> are the same, must use different dirs\n")
			os.Exit(1)
		}

		log.Printf("Comparing %v to %v", d1, d2)
		doCompare(d1, d2)
	}

	if action == "livecompare" {
		roxAPIEndpoint = strings.TrimSpace(roxAPIEndpoint)
		if roxAPIEndpoint == "" {
			panic("ROX_ENDPOINT empty")
		}

		roxAPIToken = strings.TrimSpace(roxAPIToken)
		if roxAPIToken == "" {
			panic("ROX_API_TOKEN empty")
		}

		roxAPIEndpointR = strings.TrimSpace(roxAPIEndpointR)
		if roxAPIEndpointR == "" {
			panic("ROX_ENDPOINTR empty")
		}

		roxAPITokenR = strings.TrimSpace(roxAPITokenR)
		if roxAPITokenR == "" {
			panic("ROX_API_TOKENR empty")
		}

		destDirL := filepath.Join(dumpBaseDir, "left")
		check(os.RemoveAll(destDirL))
		log.Printf("Dumping files into: %s", destDirL)
		dumpAllScannerV2Scans(destDirL, roxAPIEndpoint, roxAPIToken)

		destDirR := filepath.Join(dumpBaseDir, "right")
		check(os.RemoveAll(destDirR))
		log.Printf("Dumping files into: %s", destDirR)
		dumpAllScannerV2Scans(destDirR, roxAPIEndpointR, roxAPITokenR)

		log.Printf("Comparing %v to %v", destDirL, destDirR)
		doCompare(destDirL, destDirR)
	}
}

type apiImages struct {
	Images []*apiImage `json:"images"`
}

type apiImage struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Components  int    `json:"components"`
	CVEs        int    `json:"cves"`
	FixableCVEs int    `json:"fixableCves"`
}

func getReq(url, token string) *http.Request {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	check(err)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	return req
}

func dumpAllScannerV2Scans(destDir, endpoint, token string) {
	// create an 'index' file that lists all images scanned
	// call the ACS API to list all the images currently scanned
	// sort that list so that image

	// create files names based on digest of image
	// data/2024-11-16-HHMMSS/12345673423456234512345234523452345234523452345.json

	url := fmt.Sprintf("https://%v/v1/images", endpoint)
	req := getReq(url, token)

	resp, err := http.DefaultClient.Do(req)
	check(err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	check(err)

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("unexpected status code: %v %q %q", resp.StatusCode, resp.Status, respBody)
	}

	pretty := &bytes.Buffer{}
	err = json.Indent(pretty, respBody, "", "  ")
	check(err)

	fpath := filepath.Join(destDir, indexFileNmae)

	os.MkdirAll(filepath.Dir(fpath), 0777)
	err = os.WriteFile(fpath, pretty.Bytes(), 0777)
	check(err)

	tmp := &apiImages{}
	err = json.Unmarshal(respBody, tmp)
	check(err)

	sort.Slice(tmp.Images, func(i, j int) bool {
		ii := tmp.Images[i]
		jj := tmp.Images[j]
		return ii.Name < jj.Name
	})

	for _, img := range tmp.Images {
		log.Printf("Dumping: %s\n", img.Name)
		dumpImage(destDir, img.ID, endpoint, token)
	}
}

func dumpImage(destDir, id, endpoint, token string) {
	url := fmt.Sprintf("https://%v/v1/images/%v", endpoint, id)
	req := getReq(url, token)

	resp, err := http.DefaultClient.Do(req)
	check(err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	check(err)

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("unexpected status code: %v %q %q", resp.StatusCode, resp.Status, respBody)
	}

	pretty := &bytes.Buffer{}
	err = json.Indent(pretty, respBody, "", "  ")
	check(err)

	id = strings.TrimPrefix(id, "sha256:")
	fpath := filepath.Join(destDir, fmt.Sprintf("%v.json", id))

	os.MkdirAll(filepath.Dir(fpath), 0777)
	err = os.WriteFile(fpath, pretty.Bytes(), 0777)
	check(err)
}

func doCompare(d1, d2 string) {
	dataB, err := os.ReadFile(filepath.Join(d1, indexFileNmae))
	check(err)

	lIndex := &apiImages{}
	err = json.Unmarshal(dataB, lIndex)
	check(err)

	dataB, err = os.ReadFile(filepath.Join(d2, indexFileNmae))
	check(err)

	rIndex := &apiImages{}
	err = json.Unmarshal(dataB, rIndex)
	check(err)

	lMap := toMap(lIndex.Images)
	rMap := toMap(rIndex.Images)

	dels := []string{}
	for id, img := range lMap {
		if _, ok := rMap[id]; !ok {
			log.Printf("image %q (%v) not in %v", id, img.Name, d2)
			dels = append(dels, id)
		}
	}

	leftOnly := len(dels)
	for _, del := range dels {
		delete(lMap, del)
	}

	dels = []string{}
	for id, img := range rMap {
		if _, ok := lMap[id]; !ok {
			log.Printf("image %q (%v) not in %v", id, img.Name, d1)
			dels = append(dels, id)
		}
	}

	rightOnly := len(dels)
	for _, del := range dels {
		delete(rMap, del)
	}

	// Each map now has identical entries
	identical := 0
	for id, limg := range lMap {
		rimg := rMap[id]

		if deepCompare(d1, d2, limg, rimg) {
			identical++
		}
	}

	log.Printf("%d image only in  left dir %q", leftOnly, d1)
	log.Printf("%d image only in right dir %q", rightOnly, d2)
	log.Printf("%d/%d images in common are identical", identical, len(lMap))
}

func toMap(images []*apiImage) map[string]*apiImage {
	m := map[string]*apiImage{}
	for _, image := range images {
		m[image.ID] = image
	}

	return m
}

func deepCompare(d1 string, d2 string, l *apiImage, r *apiImage) bool {
	printed := false
	printFn := func() {
		if !printed {
			log.Printf("Comparing image: %v (%v)", l.ID, r.Name)
			printed = true
		}
	}

	if l.Components != r.Components {
		printFn()
		log.Printf("  Component counts differ (%v vs %v)", l.Components, r.Components)
	}
	if l.CVEs != r.CVEs {
		printFn()
		log.Printf("  CVE counts differ (%v vs %v)", l.CVEs, r.CVEs)
	}

	if l.FixableCVEs != r.FixableCVEs {
		printFn()
		log.Printf("  FixableCVE counts differ (%v vs %v)", l.FixableCVEs, r.FixableCVEs)
	}

	limg, lcomps := loadACSImage(d1, l)
	rimg, rcomps := loadACSImage(d2, r)

	if limg.Scan == nil && rimg.Scan != nil {
		printFn()
		log.Printf("  Scan missing from %q", d1)
		return false
	} else if limg.Scan != nil && rimg.Scan == nil {
		printFn()
		log.Printf("  Scan missing from %q", d2)
		return false
	} else if limg.Scan == nil && rimg.Scan == nil {
		return false
	}

	lscan := limg.Scan
	rscan := rimg.Scan

	if len(lscan.Components) != len(rscan.Components) {
		printFn()
		log.Printf("  Deep Component counts differ (%v vs %v)", len(lscan.Components), len(rscan.Components))
	}

	dels := []string{}
	for namever := range lcomps {
		if _, ok := rcomps[namever]; !ok {
			printFn()
			log.Printf("  Deep comp name ver %q not in %v", namever, d2)
			dels = append(dels, namever)
		}
	}

	for _, del := range dels {
		delete(lcomps, del)
	}

	dels = []string{}
	for namever := range rcomps {
		if _, ok := lcomps[namever]; !ok {
			printFn()
			log.Printf("  Deep comp name ver %q not in %v", namever, d1)
			dels = append(dels, namever)
		}
	}

	for _, del := range dels {
		delete(rcomps, del)
	}

	for namever, lc := range lcomps {
		rc := rcomps[namever]

		if len(lc.Vulns) != len(rc.Vulns) {
			printFn()
			log.Printf("  Vuln counts differ for %q (%d vs %d)", namever, len(lc.Vulns), len(rc.Vulns))
		}

		for cve, lvuln := range lc.Vulns {
			rvuln, ok := rc.Vulns[cve]
			if !ok {
				printFn()
				log.Printf("  CVE %q missing for %q in %v ", cve, namever, d2)
				continue
			}

			if lvuln.CVSS != rvuln.CVSS {
				printFn()
				log.Printf("  CVSS %q differs for %q (%v != %v) ", cve, namever, lvuln.CVSS, rvuln.CVSS)
			}

			if lvuln.Link != rvuln.Link {
				printFn()
				log.Printf("  Link %q differs for %q (%v != %v) ", cve, namever, lvuln.Link, rvuln.Link)
			}

			if lvuln.Severity != rvuln.Severity {
				printFn()
				log.Printf("  Severity %q differs for %q (%v != %v) ", cve, namever, lvuln.Severity, rvuln.Severity)
			}

			if lvuln.Summary != rvuln.Summary {
				printFn()
				log.Printf("  Summary %q differs for %q (%v != %v) ", cve, namever, lvuln.Summary, rvuln.Summary)
			}
		}

		for cve, rvuln := range rc.Vulns {
			lvuln, ok := lc.Vulns[cve]
			if !ok {
				printFn()
				log.Printf("  CVE %q missing for %q in %v ", cve, namever, d1)
				continue
			}

			if lvuln.CVSS != rvuln.CVSS {
				printFn()
				log.Printf("  CVSS %q differs for %q (%v != %v) ", cve, namever, lvuln.CVSS, rvuln.CVSS)
			}

			if lvuln.Link != rvuln.Link {
				printFn()
				log.Printf("  Link %q differs for %q (%v != %v) ", cve, namever, lvuln.Link, rvuln.Link)
			}

			if lvuln.Severity != rvuln.Severity {
				printFn()
				log.Printf("  Severity %q differs for %q (%v != %v) ", cve, namever, lvuln.Severity, rvuln.Severity)
			}

			if lvuln.Summary != rvuln.Summary {
				printFn()
				log.Printf("  Summary %q differs for %q (%v != %v) ", cve, namever, lvuln.Summary, rvuln.Summary)
			}
		}
	}

	// !printed indicates the images are identical
	return !printed

}

type acsImage struct {
	Scan *acsScan `json:"scan"`
}

type acsScan struct {
	Components []*acsComponent `json:"components"`
}

type acsComponent struct {
	Name    string     `json:"name"`
	Version string     `json:"version"`
	Vulns   []*acsVuln `json:"vulns"`
}

type acsVuln struct {
	CVE      string  `json:"cve"`
	Summary  string  `json:"summary"`
	Link     string  `json:"link"`
	CVSS     float32 `json:"cvss"`
	Severity string  `json:"severity"`
}

type compareComponent struct {
	Name    string
	Version string
	Vulns   map[string]*acsVuln
}

func loadACSImage(dir string, apiImage *apiImage) (*acsImage, map[string]*compareComponent) {
	id := strings.TrimPrefix(apiImage.ID, "sha256:")
	fpath := filepath.Join(dir, fmt.Sprintf("%v.json", id))

	dataB, err := os.ReadFile(fpath)
	check(err)

	img := &acsImage{}
	err = json.Unmarshal(dataB, img)
	check(err)

	comps := map[string]*compareComponent{}
	if img.Scan == nil {
		return img, comps
	}

	for _, c := range img.Scan.Components {
		id := fmt.Sprintf("%s/%s", c.Name, c.Version)

		vulns := map[string]*acsVuln{}
		for _, v := range c.Vulns {
			vulns[v.CVE] = v
		}
		comps[id] = &compareComponent{
			Name:    c.Name,
			Version: c.Version,
			Vulns:   vulns,
		}

	}

	return img, comps
}
