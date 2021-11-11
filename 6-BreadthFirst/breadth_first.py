from collections import deque


def person_is_seller(name: str) -> bool:
    return name[-1] == 'm'


def search(name: str) -> bool:
    graph: dict = {"you": ["alice", "bob", "claire"], "bob": ["anuj", "peggy"], "alice": ["peggy"],
                   "claire": ["thom", "jonny"], "anuj": [], "peggy": [], "thom": [], "jonny": []}
    search_queue: deque = deque()
    search_queue += graph[name]
    searched: list = []
    while search_queue:
        person = search_queue.popleft()
        if person not in searched:
            if person_is_seller(person):
                print(person + " is а mango seller!")
                return True
            else:
                search_queue += graph[person]
                searched.append(person)
    return False


def main():
    search("you")


if __name__ == '__main__':
    main()
