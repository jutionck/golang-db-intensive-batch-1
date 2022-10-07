package utils

const (
	// customer
	INSERT_CUSTOMER      = `insert into customer(name,address,job) values (:name,:address,:job)`
	UPDATE_CUSTOMER      = `update customer set name=:name, address=:address, job=:job where id=:id`
	DELETE_CUSTOMER      = `delete from customer where id=$1`
	SELECT_CUSTOMER_LIST = `select id,name,address,job from customer limit $1 offset $2`
	SELECT_CUSTOMER_ID   = `select id,name,address,job from customer where id=$1`
)
