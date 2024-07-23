---
title: "Dates with Swagger and Spring Boot"
date: "2020-04-02"
summary: "Handle dates in swagger-codegen controllers"
description: "Handle dates in swagger-codegen controllers"
toc: true
readTime: true
autonumber: true
math: true
tags: ["programming","java"]
showTags: false
---

## Problem

When using a date parameter in a API call with a Swagger generated Controller interface I noticed some problems
concerning the date format.
 
I was struggling to find a solution to receive the dates without adding the
`@DateTimeFormat(iso = DateTimeFormat.ISO.DATE_TIME)` annotation to the date parameter, which defeats the purpose of using generated code.

## Solution

The solution I found [here](https://github.com/swagger-api/swagger-codegen/issues/4113) is to add a formatter
to a `WebMvcConfigurer` class.

**Example:**

```java
@Configuration
public class RestConfiguration extends WebMvcConfigurerAdapter {

    @Override
    public void addFormatters(FormatterRegistry registry) {
        DateTimeFormatterRegistrar registrar = new DateTimeFormatterRegistrar();
        registrar.setUseIsoFormat(true);
        registrar.registerFormatters(registry);
    }
}
```
