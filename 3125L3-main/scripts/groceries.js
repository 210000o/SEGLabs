	
// Array of products, each product is an object with different fieldset
// A set of ingredients should be added to products		 

var products = [
	{
		name: "brocoli",
		vegetarian: true,
		glutenFree: true,
		meat: false,
		price: 1.99
	},
	{
		name: "onion",
		vegetarian: true,
		glutenFree: true,
		meat: false,
		price: 1.54
	},
	{
		name: "carrot",
		vegetarian: true,
		glutenFree: true,
		meat: false,
		price: 2.5
	},
	{
		name: "cabbage",
		vegetarian: true,
		glutenFree: true,
		meat: false,
		price: 1.44
	},
	{
		name: "bread",
		vegetarian: true,
		glutenFree: false,
		meat: false,
		price: 2.35
	},
	{
		name: "chicken chest",
		vegetarian: false,
		glutenFree: true,
		meat: true,
		price: 5.99
	},
	{
		name: "beef balls",
		vegetarian: false,
		glutenFree: true,
		meat: true,
		price: 7.56
	},
	{
		name: "potato",
		vegetarian: true,
		glutenFree: true,
		meat: false,
		price: 2.04
	},
	{
		name: "pork",
		vegetarian: false,
		glutenFree: true,
		meat: true,
		price: 1.99
	},
	{
		name: "salmon",
		vegetarian: false,
		glutenFree: true,
		meat: true,
		price: 10.00
	}
];
	


// given restrictions provided, make a reduced list of products
// prices should be included in this list, as well as a sort based on price

function restrictListProducts(prods, restriction) {
	let product_names = [];
	for (let i=0; i<prods.length; i+=1) {
		if ((restriction == "Vegetarian") && (prods[i].vegetarian == true)){
			product_names.push(prods[i].name);
		}
		else if ((restriction == "GlutenFree") && (prods[i].glutenFree == true)){
			product_names.push(prods[i].name);
		}
		else if ((restriction == "Meat") && (prods[i].meat == true)){
			product_names.push(prods[i].name);
		}
		else if (restriction == "None"){
			product_names.push(prods[i].name);
		}
	}
	return product_names;
}

// Calculate the total price of items, with received parameter being a list of products
function getTotalPrice(chosenProducts) {
	totalPrice = 0;
	for (let i=0; i<products.length; i+=1) {
		if (chosenProducts.indexOf(products[i].name) > -1){
			totalPrice += products[i].price;
		}
	}
	return totalPrice;
}
