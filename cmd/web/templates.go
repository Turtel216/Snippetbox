package main

import (
  "html/template"
  "path/filepath"

  "snippetbox.dimitrios_papakonstantinou.com/internal/models"
)

type templateData struct {
  Snippet *models.Snippet
  Snippets []*models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
  // Initialize a new map to act as the cache.
  cache := map[string]*template.Template{}

  pages, err := filepath.Glob("./ui/html/pages/*.html")
  if err != nil {
    return nil, err
  }

  for _, page := range pages {
    name := filepath.Base(page)

    files := []string {
      "./ui/html/base.html",
      "./ui/html/partials/nav.html",
      page,
    }

    ts, err := template.ParseFiles(files ...)
    if err != nil {
      return nil, err
    }

    cache[name] = ts
  }

  return cache, nil
}