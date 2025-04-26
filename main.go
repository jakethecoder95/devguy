package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

// loggingMiddleware is a middleware that logs incoming HTTP requests.
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        // Capture the response status and size
        lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
        next.ServeHTTP(lrw, r)

        duration := time.Since(start)

        // Extract client IP
        clientIP := getClientIP(r)

        // Log the details
        log.Printf(
            "%s - [%s] \"%s %s %s\" %d %d \"%s\" \"%s\" %v",
            clientIP,
            start.Format(time.RFC1123),
            r.Method,
            r.RequestURI,
            r.Proto,
            lrw.statusCode,
            lrw.size,
            r.Referer(),
            r.UserAgent(),
            duration,
        )
    })
}

// loggingResponseWriter wraps http.ResponseWriter to capture status code and response size.
type loggingResponseWriter struct {
    http.ResponseWriter
    statusCode int
    size       int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
    lrw.statusCode = code
    lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
    size, err := lrw.ResponseWriter.Write(b)
    lrw.size += size
    return size, err
}

// getClientIP extracts the client's real IP address from the request, considering proxy headers.
func getClientIP(r *http.Request) string {
    // Check X-Forwarded-For header (comma-separated list of IPs)
    xForwardedFor := r.Header.Get("X-Forwarded-For")
    if xForwardedFor != "" {
        // The first IP in the list is the client's IP
        ips := strings.Split(xForwardedFor, ",")
        if len(ips) > 0 {
            ip := strings.TrimSpace(ips[0])
            if ip != "" {
                return ip
            }
        }
    }

    // Check X-Real-IP header
    xRealIP := r.Header.Get("X-Real-IP")
    if xRealIP != "" {
        return strings.TrimSpace(xRealIP)
    }

    // Fallback to RemoteAddr
    ip, _, err := net.SplitHostPort(r.RemoteAddr)
    if err != nil {
        return r.RemoteAddr // Return as-is if unable to split
    }
    return ip
}

// helloHandler is a simple HTTP handler that responds with a greeting.
func helloHandler(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./index.html");
    fmt.Fprintf(w, "Hello, World! You've requested: %s\n", r.URL.Path)
}

func main() {
    // Create a new ServeMux and register handlers
    mux := http.NewServeMux()
    mux.HandleFunc("/", helloHandler)

    // Wrap the ServeMux with the logging middleware
    loggedMux := loggingMiddleware(mux)

    // Define the server
    server := &http.Server{
        Addr:         ":8080", // Listen on port 8080
        Handler:      loggedMux,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  120 * time.Second,
    }

    log.Printf("Starting server on %s", server.Addr)
    if err := server.ListenAndServe(); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
