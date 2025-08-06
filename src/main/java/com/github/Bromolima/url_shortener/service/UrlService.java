package com.github.Bromolima.url_shortener.service;

import com.github.Bromolima.url_shortener.entities.Url;
import com.github.Bromolima.url_shortener.repository.UrlRepository;
import jakarta.servlet.http.HttpServletRequest;
import org.apache.commons.lang3.RandomStringUtils;
import org.springframework.stereotype.Service;

import java.time.LocalDateTime;

@Service
public class UrlService {
    private final UrlRepository urlRepository;

    public UrlService(UrlRepository urlRepository) {
        this.urlRepository = urlRepository;
    }

    public Url shortenUrl(String longUrl, LocalDateTime expiresAt) {
        String id;
        do {
            id = RandomStringUtils.randomAlphanumeric(6, 7);
        } while (urlRepository.existsById(id));

        urlRepository.save(new Url(id, longUrl, expiresAt));

        return urlRepository.findById(id).get();
    }
}
