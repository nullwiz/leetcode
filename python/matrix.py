import numpy as np
import unittest


def createSquareMatrix():
    rand = np.random.randint(0, 10, size=(4, 4))
    return rand


def createNonSquareMatrix():
    rand = np.random.randint(0, 10, size=(4, 2))
    return rand


def multiply(m1, m2):
    if len(m1[0]) != len(m2):
        return 0
    ret = [[0 for _ in m2[0]] for _ in m1]
    for i in range(len(m2)):
        for j in range(len(m2[0])):
            for z in range(len(m1[0])):
                ret[i][j] += m1[i][z] * m2[z][j]
    pprint(ret)
    return ret


def msum(m1, m2):
    if len(m1[0]) != len(m2[0]) or len(m1) != len(m2):
        return 0
    ret = [[0 for _ in m2[0]] for _ in m1]
    for i in range(len(m1)):
        for j in range(len(m1[0])):
            ret[i][j] = m1[i][j] + m2[i][j]
    pprint(ret)
    return ret


def direct_sum(m1, m2):
    rows = len(m1) + len(m2)
    cols = len(m1[0]) + len(m2[0])
    ret = [[0 for _ in range(cols)] for _ in range(rows)]
    for i in range(len(m1)):
        for j in range(len(m1[0])):
            ret[i][j] = m1[i][j]

    for i in range(len(m1), rows):
        for j in range(len(m1[0]), cols):
            ret[i][j] = m2[i - len(m1)][j - len(m1[0])]
    pprint(ret)
    return ret


def pprint(matrix):
    for row in matrix:
        print(row)


def flipMatrix(matrix):
    matrix[:, 2] = matrix[::-1, 2]
    matrix[0, :] = matrix[0, ::-1]
    print(matrix)


class MatrixOperationsTest(unittest.TestCase):
    def setUp(self):
        self.square_matrix_1 = [[1, 2], [3, 4]]
        self.square_matrix_2 = [[5, 6], [7, 8]]
        self.non_square_matrix_1 = [[1, 2], [3, 4], [5, 6]]
        self.non_square_matrix_2 = [[7, 8], [9, 10], [11, 12]]

    def test_multiply(self):
        np_result = np.matmul(self.square_matrix_1, self.square_matrix_2)
        result = np.array(multiply(self.square_matrix_1, self.square_matrix_2))
        self.assertTrue(np.array_equal(result, np_result))

    def test_msum(self):
        result = np.sum([self.square_matrix_1, self.square_matrix_2], axis=0)
        self.assertTrue(
            np.array_equal(
                np.array(msum(self.square_matrix_1, self.square_matrix_2)),
                result,
            )
        )

    def test_direct_sum(self):
        np_result = np.array(
            [
                [1, 2, 0, 0],
                [3, 4, 0, 0],
                [5, 6, 0, 0],
                [0, 0, 7, 8],
                [0, 0, 9, 10],
                [0, 0, 11, 12],
            ]
        )
        result = np.array(
            direct_sum(self.non_square_matrix_1, self.non_square_matrix_2)
        )

        self.assertTrue(np.array_equal(result, np_result))


if __name__ == "__main__":
    unittest.main()
