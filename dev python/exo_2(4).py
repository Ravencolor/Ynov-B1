class StringUtils:
    new_line = "\n"
    tab = "\t"

    @staticmethod
    def get_first(str):
        return str[0] if str else None

    @staticmethod
    def get_last(str):
        return str[-1] if str else None

    @staticmethod
    def get_substring(str, first, last):
        return str[first:last] if str and first < last else None
    
    # Utilisation des propriétés statiques
print(StringUtils.new_line)
print(StringUtils.tab)

# Appel des méthodes statiques
print(StringUtils.get_first("Hello"))
print(StringUtils.get_last("Hello"))
print(StringUtils.get_substring("Hello", 1, 4))