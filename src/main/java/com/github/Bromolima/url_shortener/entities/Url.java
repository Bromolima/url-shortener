package com.github.Bromolima.url_shortener.entities;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;
import org.springframework.data.mongodb.core.index.Indexed;

import java.time.LocalDateTime;

@Document(collection = "url")
public class Url {
    @Id
    private String id;

    private String longUrl;

    @Indexed(expireAfter = "0s")
    private LocalDateTime expiresAt;

    public Url() { }

    public Url(String id, String longUrl, LocalDateTime expiresAt) {
        this.id = id;
        this.longUrl = longUrl;
        this.expiresAt = expiresAt;
    }

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getLongUrl() {
        return longUrl;
    }

    public void setLongUrl(String longUrl) {
        this.longUrl = longUrl;
    }

    public LocalDateTime getExpiresAt() {
        return expiresAt;
    }

    public void setExpiresAt(LocalDateTime expiresAt) {
        this.expiresAt = expiresAt;
    }
}
