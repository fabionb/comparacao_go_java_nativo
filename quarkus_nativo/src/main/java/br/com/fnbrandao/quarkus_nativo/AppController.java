package br.com.fnbrandao.quarkus_nativo;

import javax.inject.Inject;
import javax.persistence.EntityManager;
import javax.transaction.Transactional;
import javax.ws.rs.GET;
import javax.ws.rs.Path;

@Path("/")
public class AppController {

	@Inject
	EntityManager entityManager;

	@GET
	@Transactional
	public void index() {
		Test1Entity e = new Test1Entity();
		entityManager.persist(e);
	}

}
