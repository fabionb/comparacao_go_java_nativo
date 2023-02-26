package br.com.fnbrandao.springboot_nativo;

import jakarta.persistence.EntityManager;
import jakarta.transaction.Transactional;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class AppController {

	@Autowired
	private EntityManager entityManager;

	@GetMapping("/")
	@Transactional
	public void index() {
		Test1Entity e = new Test1Entity();
		entityManager.persist(e);
	}

}
