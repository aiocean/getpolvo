package server

import (
	"log"
	"net/http"
	"os"

	"github.com/google/wire"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	v1 "pkg.aiocean.dev/polvogo/aiocean/polvo/v1"
)


var WireSet = wire.NewSet(
	NewServer,
)

type Server struct {
	logger *zap.Logger
	polvoClient v1.PolvoServiceClient
}

func NewServer(logger *zap.Logger, polvoClient v1.PolvoServiceClient) *Server {
	return &Server{
		logger: logger,
		polvoClient: polvoClient,
	}
}

func (s *Server) handler(w http.ResponseWriter, r *http.Request) {
	packageName, versionName := parseVersionOrn(r.RequestURI)

	if packageName == "" || versionName == "" {
		http.Error(w, "package name or version name is missing", http.StatusBadRequest)
		return
	}

	request := &v1.GetVersionRequest{
		Orn: "packages/" + packageName + "/versions/" + versionName,
	}

	response, err := s.polvoClient.GetVersion(r.Context(), request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if response.GetVersion().GetManifestUrl() == "" {
		http.Error(w, "Manifest url is missing", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, response.GetVersion().GetManifestUrl(), http.StatusTemporaryRedirect)
	return
}


func (s *Server) Serve () {
	log.Print("starting server...")
	http.HandleFunc("/", s.handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	h2s := &http2.Server{}
	handler := http.HandlerFunc(s.handler)

	server := &http.Server{
		Addr:    ":"+port,
		Handler: h2c.NewHandler(handler, h2s),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
