/**
 * Fonction pour gérer l'événement DOMContentLoaded et afficher les détails de confirmation de commande.
 * @event DOMContentLoaded
 */
document.addEventListener('DOMContentLoaded', () => {
    /**
     * Récupérer les détails de la commande depuis le stockage local.
     * @type {Object}
     */
    const order = JSON.parse(localStorage.getItem('order'));

    if (order) {
        /**
         * Élément pour afficher l'ID de commande.
         * @type {HTMLElement}
         */
        const orderIdElement = document.getElementById('order-id');

        /**
         * Élément pour afficher le prix total de la commande.
         * @type {HTMLElement}
         */
        const totalPriceElement = document.getElementById('total-price');

        /**
         * Élément conteneur pour contenir les détails de confirmation.
         * @type {HTMLElement}
         */
        const confirmationDetails = document.createElement('div');

        // Définir l'ID de commande et le prix total dans les éléments respectifs
        orderIdElement.textContent = order.orderId;
        totalPriceElement.textContent = (order.totalAmount / 100).toFixed(2);

        /**
         * Élément conteneur pour contenir les détails de contact.
         * @type {HTMLElement}
         */
        const contactDetails = document.createElement('div');
        contactDetails.innerHTML = `
            <h3>Détails de contact :</h3>
            <p><strong>Nom :</strong> ${order.contact.lastName}</p>
            <p><strong>Prénom :</strong> ${order.contact.firstName}</p>
            <p><strong>Adresse :</strong> ${order.contact.address}</p>
            <p><strong>Ville :</strong> ${order.contact.city}</p>
            <p><strong>Email :</strong> ${order.contact.email}</p>
        `;

        /**
         * Élément conteneur pour contenir la liste des produits commandés.
         * @type {HTMLElement}
         */
        const productsList = document.createElement('div');
        productsList.innerHTML = `
            <h3>Produits commandés :</h3>
            <ul>
                ${order.products.map(product => `
                    <li>
                        <strong>${product.name}</strong> - Objectif: ${product.lens} - Prix unitaire: ${(product.price / 100).toFixed(2)} €
                    </li>
                `).join('')}
            </ul>
        `;

        // Ajouter les détails de contact et la liste des produits au conteneur de détails de confirmation
        confirmationDetails.appendChild(contactDetails);
        confirmationDetails.appendChild(productsList);

        // Ajouter le conteneur de détails de confirmation au corps de la page
        document.body.appendChild(confirmationDetails);

        // Supprimer les détails de commande du stockage local
        localStorage.removeItem('order');
    } else {
        console.error('Aucune commande trouvée.');
        // Rediriger vers la page d'accueil si aucune commande n'est trouvée
        window.location.href = '/';
    }
});
