$(document).ready(function() {
    let editIndex = null;
    let contacts = [];

    // Charger les contacts depuis le local storage
    function chargerContacts() {
        contacts = JSON.parse(localStorage.getItem('contacts')) || [];
        $('#liste-utilisateur').empty();
        contacts.forEach((contact, index) => {
            ajouterContactHTML(contact, index);
        });
        desactiverBoutonsAlphabet();
    }

    // Trier les contacts par ordre alphabétique
    function trierContacts() {
        contacts.sort((a, b) => a.nom.localeCompare(b.nom));
    }

    // Afficher les informations d'un utilisateur
    function afficherinfo() {
        $(this).siblings('.affichage').toggle();
    }

    // Afficher le formulaire pour ajouter un nouveau contact
    function afficherform() {
        $('#ajout-contact').toggle();
        // Réinitialiser le formulaire
        $('#civil').val('monsieur');
        $('#prenom').val('');
        $('#nom').val('');
        $('#num').val('');
        editIndex = null;
    }

    // Ajouter ou modifier un contact
    function ajouterContact() {
        let civil = $('#civil').val();
        let prenom = $('#prenom').val();
        let nom = $('#nom').val();
        let num = $('#num').val();

        if (prenom && nom && num) {
            let contact = {
                civil: civil,
                prenom: prenom,
                nom: nom,
                num: num
            };

            if (editIndex !== null) {
                // Modifier le contact existant
                contacts[editIndex] = contact;
                editIndex = null;
            } else {
                // Ajouter un nouveau contact
                contacts.push(contact);
            }

            trierContacts();
            localStorage.setItem('contacts', JSON.stringify(contacts));
            chargerContacts();
            $('#ajout-contact').hide();
        } else {
            alert('Veuillez remplir tous les champs.');
        }
    }

    function ajouterContactHTML(contact, index) {
        let utilisateur = `<li>
        <p class="nom-utilisateur">${contact.nom} ${contact.prenom}</p>
        <div class="affichage" style="display: none;">
            <p class="nom-complet">${contact.civil} ${contact.prenom} ${contact.nom}</p>
            <p class="tel-utilisateur">${contact.num}</p>
            <p class="modification" data-index="${index}"><i class="fas fa-pen"></i> Editer ce contact</p>
            <p class="suppression" data-index="${index}"><i class="fas fa-trash"></i> Supprimer ce contact</p>
        </div>
        </li>`;

        $('#liste-utilisateur').append(utilisateur);
    }

    // Afficher le formulaire de modification prérempli
    function modifierContact() {
        editIndex = $(this).data('index');
        let contact = contacts[editIndex];

        $('#civil').val(contact.civil);
        $('#prenom').val(contact.prenom);
        $('#nom').val(contact.nom);
        $('#num').val(contact.num);
        $('#ajout-contact').show();
    }

    // Supprimer un contact spécifique
    function supprimerContact() {
        let index = $(this).data('index');
        contacts.splice(index, 1);
        localStorage.setItem('contacts', JSON.stringify(contacts));
        chargerContacts();
    }

    // Supprimer tous les contacts
    function supprimerTousLesContacts() {
        localStorage.removeItem('contacts');
        contacts = [];
        chargerContacts();
    }

    // Filtrer les contacts par la première lettre du nom
    function filterContacts(letter) {
        $('.nom-utilisateur').each(function() {
            let nom = $(this).text();
            if (nom.charAt(0).toUpperCase() === letter) {
                $(this).parent().show();
            } else {
                $(this).parent().hide();
            }
        });
    }

    // Désactiver les boutons de l'alphabet sans contacts correspondants
    function desactiverBoutonsAlphabet() {
        let alphabet = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
        for (let i = 0; i < alphabet.length; i++) {
            let letter = alphabet[i];
            let hasContacts = contacts.some(contact => contact.nom.charAt(0).toUpperCase() === letter);
            let button = $('.alpha-btn').eq(i);
            if (!hasContacts) {
                button.prop('disabled', true);
            } else {
                button.prop('disabled', false);
            }
        }
    }

    // Attacher les événements
    $(document).on('click', '.nom-utilisateur', afficherinfo);
    $('#ajout-utilisateur1').on('click', afficherform);
    $('#ajout-utilisateur2').on('click', ajouterContact);
    $('#supprimer-tous').on('click', supprimerTousLesContacts);
    $(document).on('click', '.modification', modifierContact);
    $(document).on('click', '.suppression', supprimerContact);

    // Attacher un événement aux boutons de l'alphabet
    $('.alpha-btn').on('click', function() {
        let letter = $(this).text().toUpperCase();
        filterContacts(letter);
    });

    // Charger les contacts au démarrage
    chargerContacts();
});
