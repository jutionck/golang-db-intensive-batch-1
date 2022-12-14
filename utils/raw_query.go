package utils

const (
	// customer
	INSERT_CUSTOMER      = `insert into customer(name,address,job) values ($1,$2,$3)`
	UPDATE_CUSTOMER      = `update customer set name=$1, address=$2, job=$3 where id=$4`
	DELETE_CUSTOMER      = `delete from customer where id=$1`
	SELECT_CUSTOMER_LIST = `select id,name,address,job from customer limit $1 offset $2`
	SELECT_CUSTOMER_ID   = `select id,name,address,job from customer where id=$1`
)
