from itertools import permutations

def castle(layer, used_numbers, solution):

    if layer == 5:
        return solution

    layersize = [5, 4, 3, 2, 1]
    currentsize = layersize[layer]

    if layer == 0:
        for first_layer in permutations(range(1, 16), currentsize):
            result = castle(layer + 1, set(first_layer), [list(first_layer)])
            if result:
                return result
        return None

    oldlayer = solution[-1]
    for nextlayer in permutations(range(1, 16), currentsize):
        valid = True
        for d in range(currentsize):
            if abs(oldlayer[d] - oldlayer[d + 1]) != nextlayer[d]:
                valid = False
                break
            if nextlayer[d] in used_numbers:
                valid = False
                break
        if valid:
            result = castle(
                layer + 1, used_numbers | set(nextlayer), solution + [list(nextlayer)]
            )
            if result:
                return result

    return None

solution = castle(0, set(), [])
if solution:
    for layer in solution:
        print(layer)
else:
    print("no solution")