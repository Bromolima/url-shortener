package com.github.Bromolima.url_shortener.controller;

import com.github.Bromolima.url_shortener.controller.dto.ShortenUrlRequest;
import com.github.Bromolima.url_shortener.controller.dto.ShortenUrlResponse;
import com.github.Bromolima.url_shortener.repository.UrlRepository;
import com.github.Bromolima.url_shortener.service.UrlService;
import jakarta.servlet.http.HttpServletRequest;
import org.springframework.http.HttpHeaders;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.net.URI;
import java.time.LocalDateTime;

@RestController
public class UrlController {
    private final UrlService urlService;
    private final UrlRepository urlRepository;

    public UrlController(UrlService urlService, UrlRepository urlRepository) {
        this.urlService = urlService;
        this.urlRepository = urlRepository;
    }

    @PostMapping(value = "/shorten-url")
    public ResponseEntity<ShortenUrlResponse> shortenUrl (@RequestBody ShortenUrlRequest request,
                                                          HttpServletRequest servletRequest) {

        var Url = urlService.shortenUrl(request.url(), LocalDateTime.now().plusMinutes(1));

        var redirectUrl = servletRequest.getRequestURL().toString().replace("shorten-url", Url.getId());

        return ResponseEntity.ok(new ShortenUrlResponse(redirectUrl));
    }

    @GetMapping("{id}")
    public ResponseEntity<Void> redirect(@PathVariable("id") String id) {
        var url = urlRepository.findById(id);

        if (url.isEmpty()) {
            return ResponseEntity.notFound().build();
        }

        HttpHeaders headers = new HttpHeaders();
        headers.setLocation(URI.create(url.get().getLongUrl()));

        return ResponseEntity.status(HttpStatus.FOUND).headers(headers).build();
    }
}
