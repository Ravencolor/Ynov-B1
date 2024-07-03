document.addEventListener("DOMContentLoaded", () => {
    const params = new URLSearchParams(window.location.search);
    const productId = params.get('id');

    // Récupérer les détails du produit depuis l'API en utilisant l'ID du produit
    fetch(`http://localhost:3000/api/cameras/${productId}`)
        .then(response => response.json())
        .then(product => {
            const productDetails = document.getElementById('product-details');

            // Former correctement l'URL de l'image
            const imageUrl = product.imageUrl.startsWith('http') 
                ? product.imageUrl 
                : `http://localhost:3000/images/${product.imageUrl}`;

            // Afficher les détails du produit dans le HTML
            productDetails.innerHTML = `
                <h2>${product.name}</h2>
                <img src="${imageUrl}" alt="${product.name}" class="product-image">
                <p>${product.description}</p>
                <p>${product.price / 100} €</p>
                <label for="lens">Choix de l'objectif:</label>
                <select id="lens">
                    ${product.lenses.map(lens => `<option value="${lens}">${lens}</option>`).join('')}
                </select>
            `;

            // Ajouter un écouteur d'événement pour le bouton "Ajouter au panier"
            document.getElementById('add-to-cart').addEventListener('click', () => {
                const selectedLens = document.getElementById('lens').value;
                const cartItem = {
                    id: product._id,
                    name: product.name,
                    price: product.price,
                    lens: selectedLens,
                };

                // Récupérer le panier depuis le stockage local ou créer un nouveau panier
                let cart = JSON.parse(localStorage.getItem('cart')) || [];
                cart.push(cartItem);
                localStorage.setItem('cart', JSON.stringify(cart));
                alert('Produit ajouté au panier');
            });

            // Ajouter un écouteur d'événement pour le bouton "Voir le panier"
            document.getElementById('view-cart').addEventListener('click', () => {
                window.location.href = 'panier.html';
            });
        })
        .catch(error => console.error('Erreur lors de la récupération du produit:', error));
});



