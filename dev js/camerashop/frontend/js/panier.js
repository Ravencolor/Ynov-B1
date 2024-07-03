document.addEventListener("DOMContentLoaded", () => {
    // Récupérer le panier depuis le localStorage ou initialiser un tableau vide
    const cart = JSON.parse(localStorage.getItem('cart')) || [];
    const cartItems = document.getElementById('cart-items');
    const totalPriceElement = document.getElementById('total-price');
    let totalPrice = 0;

    // Mettre à jour le prix total affiché à l'écran
    const updateTotalPrice = () => {
        totalPriceElement.textContent = (totalPrice / 100).toFixed(2);
    };

    // Supprimer un article du panier en fonction de son index
    const removeItemFromCart = (index) => {
        cart.splice(index, 1);
        localStorage.setItem('cart', JSON.stringify(cart));
        renderCartItems();
    };

    // Afficher les articles du panier à l'écran
    const renderCartItems = () => {
        cartItems.innerHTML = '';
        totalPrice = 0;

        cart.forEach((item, index) => {
            const cartItem = document.createElement('div');
            cartItem.className = 'cart-item';
            cartItem.innerHTML = `
                <h2>${item.name}</h2>
                <p>Objectif: ${item.lens}</p>
                <p>Prix: ${(item.price / 100).toFixed(2)} €</p>
                <button class="remove-item" data-index="${index}">Supprimer</button>
            `;
            cartItems.appendChild(cartItem);
            totalPrice += item.price;
        });

        // Ajouter un écouteur d'événement pour chaque bouton de suppression d'article
        document.querySelectorAll('.remove-item').forEach(button => {
            button.addEventListener('click', (e) => {
                const index = e.target.dataset.index;
                removeItemFromCart(index);
            });
        });

        updateTotalPrice();
    };

    // Afficher les articles du panier au chargement de la page
    renderCartItems();

    // Gérer la soumission du formulaire de commande
    document.getElementById('order-form').addEventListener('submit', (e) => {
        e.preventDefault();
        const formData = new FormData(e.target);
        const contact = Object.fromEntries(formData.entries());

        // Créer un tableau de produits à partir du panier
        const products = cart.map(item => ({ id: item.id, name: item.name, price: item.price, quantity: 1, lens: item.lens }));

        // Générer un identifiant de commande unique
        const orderId = Date.now().toString();

        // Créer un objet de commande contenant les informations de contact, les produits et le montant total
        const order = {
            orderId,
            contact,
            products,
            totalAmount: totalPrice
        };

        // Enregistrer la commande dans le localStorage
        localStorage.setItem('order', JSON.stringify(order));

        // Afficher une alerte de succès
        alert('Commande passée avec succès!');

        // Supprimer le panier du localStorage
        localStorage.removeItem('cart');

        // Rediriger vers la page de confirmation de commande
        window.location.href = `confirmation.html`;
    });
});
