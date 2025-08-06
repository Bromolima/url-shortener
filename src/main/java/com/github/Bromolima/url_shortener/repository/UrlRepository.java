package com.github.Bromolima.url_shortener.repository;

import com.github.Bromolima.url_shortener.entities.Url;
import org.springframework.data.mongodb.repository.MongoRepository;

public interface UrlRepository extends MongoRepository<Url, String> {
}
