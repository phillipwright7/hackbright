SELECT owners.owner_id, owners.first_name, owners.last_name, cars.car_id, cars.make, cars.model, sales.sale_id
FROM sales
JOIN owners ON owners.owner_id = sales.owner_id
JOIN cars ON cars.car_id = sales.car_id;