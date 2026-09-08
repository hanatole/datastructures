class ExpressionTree:
    @staticmethod
    def _is_higher(operator1, operator2):
        priorities = {"+": 1, "-": 1, "*": 2, "/": 2, "(": 3, ")": 0}

        return priorities[operator1] > priorities[operator2]

    @staticmethod
    def _tokenize(text):
        operators = set("+-*/()")
        tokens = []
        i, s = 0, len(text)
        while i < s:
            if text[i] in operators:
                tokens.append(text[i])
            elif text[i].isdigit():
                num = 0
                while i < s and text[i].isdigit():
                    num = num * 10 + int(text[i])
                    i += 1
                tokens.append(num)
                if i < s and text[i] in operators:
                    tokens.append(text[i])
            i += 1

        return tokens

    @classmethod
    def _create_node(cls, value):
        node = object.__new__(cls)
        node.value = value
        return node

    @classmethod
    def from_text(cls, text):
        tokens = cls._tokenize(text)
        operators = []
        operands = []
        i, l = 0, len(tokens)

        def add_node():
            node = cls._create_node(operators.pop())
            node.right = operands.pop()
            node.left = operands.pop()
            operands.append(node)

        while i < l:
            if isinstance(tokens[i], int):
                operands.append(cls._create_node(tokens[i]))
                i += 1
            else:
                if (
                    len(operators) == 0
                    or operators[-1] == "("
                    or cls._is_higher(tokens[i], operators[-1])
                ):
                    operators.append(tokens[i])
                    i += 1
                elif tokens[i] == ")":
                    while operators[-1] != "(":
                        add_node()
                    operators.pop()
                    i += 1
                else:
                    add_node()

        while operators:
            add_node()

        return operands[0] if operands else None

    def evaluate(self):
        if isinstance(self.value, int):
            return self.value
        left_val = self.left.evaluate()
        right_val = self.right.evaluate()
        if self.value == "+":
            return left_val + right_val
        elif self.value == "-":
            return left_val - right_val
        elif self.value == "*":
            return left_val * right_val
        elif self.value == "/":
            return left_val / right_val


def main():
    tree = ExpressionTree.from_text("31 + 2 * (8 - 2 + 5) - 64 / 2")
    print(tree.evaluate())


if __name__ == "__main__":
    main()
