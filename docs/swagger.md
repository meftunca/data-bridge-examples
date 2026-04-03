---
title: Swagger UI
aside: false
outline: false
---

# Swagger API Explorer

<div id="swagger-ui"></div>

<script setup>
import { onMounted } from 'vue'

onMounted(() => {
  const link = document.createElement('link')
  link.rel = 'stylesheet'
  link.href = 'https://cdn.jsdelivr.net/npm/swagger-ui-dist@5/swagger-ui.css'
  document.head.appendChild(link)

  const script = document.createElement('script')
  script.src = 'https://cdn.jsdelivr.net/npm/swagger-ui-dist@5/swagger-ui-bundle.js'
  script.onload = () => {
    window.SwaggerUIBundle({
      url: '/data-bridge-examples/swagger.json',
      dom_id: '#swagger-ui',
      deepLinking: true,
      presets: [window.SwaggerUIBundle.presets.apis],
      layout: 'BaseLayout',
    })
  }
  document.body.appendChild(script)
})
</script>

<style>
#swagger-ui {
  margin-top: 1rem;
}
.swagger-ui .topbar { display: none; }
</style>
