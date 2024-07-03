document.addEventListener("DOMContentLoaded", () => {
    // Attendre que le DOM soit chargé avant d'exécuter le code

    fetch('http://localhost:3000/api/cameras')
        // Effectuer une requête GET vers l'API pour récupérer les données des caméras
        .then(response => response.json())
        // Convertir la réponse en format JSON
        .then(data => {
            // Manipuler les données récupérées

            const productList = document.getElementById('product-list');
            // Récupérer l'élément avec l'ID 'product-list' dans le DOM

            data.forEach(product => {
                // Parcourir chaque produit dans les données

                const productItem = document.createElement('div');
                // Créer un élément <div> pour chaque produit
                productItem.className = 'product-item';
                // Ajouter la classe 'product-item' à l'élément <div>

                // Former correctement l'URL de l'image
                const imageUrl = product.imageUrl.startsWith('http') 
                    ? product.imageUrl 
                    : `http://localhost:3000/images/${product.imageUrl}`;

                productItem.innerHTML = `
                    <h2>${product.name}</h2>
                    <img src="${imageUrl}" alt="${product.name}">
                    <p>${product.description}</p>
                    <p>${product.price / 100} €</p>
                    <a href="article.html?id=${product._id}">Voir le produit</a>
                `;
                // Remplir l'élément <div> avec les informations du produit

                productList.appendChild(productItem);
                // Ajouter l'élément <div> au DOM, à l'intérieur de l'élément avec l'ID 'product-list'
            });
        })
        .catch(error => console.error('Error fetching products:', error));
        // Gérer les erreurs lors de la récupération des produits
});
