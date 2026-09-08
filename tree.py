from collections import deque


class Node:
    def __init__(self, v):
        self.v = v
        self.left = None
        self.right = None

    def __str__(self):
        return f"{self.v}"

    @classmethod
    def fromText1(cls, text):
        data = text.split()
        nodes = [None if x == "NULL" else cls(int(x)) for x in data]

        i, s, missing = 0, len(nodes), 0

        while i < s:
            if nodes[i]:
                j = 2 * i + 1 - missing
                if j < s:
                    nodes[i].left = nodes[j]
                if j + 1 < s:
                    nodes[i].right = nodes[j + 1]
            else:
                missing += 2
            i += 1

        return nodes[0]

    @classmethod
    def fromText(cls, text):
        values = text.split()
        if not values or values[0] == "NULL":
            return None

        root = cls(int(values[0]))
        queue = deque([root])
        i = 1

        while queue and i < len(values):
            node = queue.popleft()

            if values[i] != "NULL":
                node.left = cls(int(values[i]))
                queue.append(node.left)
            i += 1

            if i < len(values) and values[i] != "NULL":
                node.right = cls(int(values[i]))
                queue.append(node.right)
            i += 1

        return root

    def process(self):
        print(self, end=" ")

    def bfs(self):
        q = deque([self])
        while q:
            node = q.popleft()
            node.process()
            if node.left:
                q.append(node.left)
            if node.right:
                q.append(node.right)
        print()

    def preorder(self):
        self.process()
        if self.left:
            self.left.preorder()
        if self.right:
            self.right.preorder()

    def inorder(self):
        if self.left:
            self.left.inorder()
        self.process()
        if self.right:
            self.right.inorder()

    def postorder(self):
        if self.left:
            self.left.postorder()
        if self.right:
            self.right.postorder()
        self.process()


text = "8 5 9 1 NULL 2 4"
text = "5 1 4 NULL NULL 3 6"

root = Node.fromText(text)
root.bfs()
root.preorder()
print()
root.inorder()
print()
root.postorder()
print()

"""
    8
  5   9
 1   2  4

8 5 9 1 NULL 2 4
8 5 1 9 2 4
1 5 8 2 9 4
1 5 2 4 9 8
"""
