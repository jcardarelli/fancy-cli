create table [restaurants] (
	id INTEGER NOT NULL PRIMARY KEY,
	name TEXT,
	address TEXT,
	stars INTEGER
);


sqlite> insert into restaurants(name, address, stars) values('The French Laundry', '6640 Washington St, Yountville, CA 94599', 3);
sqlite> select * from restaurants;
1|The French Laundry|6640 Washington St, Yountville, CA 94599|3

