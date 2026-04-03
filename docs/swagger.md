---
title: Swagger UI
aside: false
outline: false
---

# Swagger API Explorer

<script setup>
import { useOpenapi } from 'vitepress-openapi'
import spec from '../swag/openapi.json'

useOpenapi({ spec })
</script>

<OASpec />
