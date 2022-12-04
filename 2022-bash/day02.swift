import Foundation

let fileName = "day02.txt"  //this is the file. we will write to and read from it
let currentDirectory = URL(fileURLWithPath: FileManager.default.currentDirectoryPath)
let fileURL = currentDirectory.appendingPathComponent(fileName)

enum GameState: String {
    case Lose = "X"
    case Draw = "Y"
    case Win = "Z"

    var correspondingPoints: Int {
        switch self {
        case .Lose:
            return 0
        case .Draw:
            return 3
        case .Win:
            return 6
        }
    }

    func handPairFor(opponentHand: Hand) -> Hand {
        switch (self, opponentHand) {
        case (.Win, .Paper), (.Draw, .Scissors), (.Lose, .Rock): return .Scissors
        case (.Win, .Rock), (.Draw, .Paper), (.Lose, .Scissors): return .Paper
        case (.Win, .Scissors), (.Draw, .Rock), (.Lose, .Paper): return .Rock
        }
    }
}
enum Hand: String {
    case Rock = "A"
    case Paper = "B"
    case Scissors = "C"

    var correspondingPoints: Int {
        switch self {
        case .Rock:
            return 1
        case .Paper:
            return 2
        case .Scissors:
            return 3
        }
    }

    func playVersus(opponentHand: Hand) -> Int {
        switch (self, opponentHand) {
        case (.Scissors, .Rock), (.Rock, .Paper), (.Paper, .Scissors):
            return 0
        case (.Rock, .Rock), (.Paper, .Paper), (.Scissors, .Scissors):
            return 3
        case (.Rock, .Scissors), (.Scissors, .Paper), (.Paper, .Rock):
            return 6
        }
    }
}

func replaceXYZWithABC(_ string: String) -> String {
    var result = string
    result = result.replacingOccurrences(of: "X", with: "A")
    result = result.replacingOccurrences(of: "Y", with: "B")
    result = result.replacingOccurrences(of: "Z", with: "C")

    return result
}

// part1
var totalPoints = 0

do {
    let text = try String(contentsOf: fileURL, encoding: .utf8)

    for line in text.split(separator: "\n") {
        let formattedLine = replaceXYZWithABC(String(line))
        let hands = formattedLine.split(separator: " ")
        let (opponent, ours) = (
            Hand(rawValue: String(hands.first!))!, Hand(rawValue: String(hands.last!))!
        )

        let roundPoints = ours.playVersus(opponentHand: opponent)
        totalPoints += roundPoints + ours.correspondingPoints
    }
} catch { print("Error: \(fileURL) \n \(error)") }

print("Total points for part 1: \(totalPoints)")

// part 2
totalPoints = 0
do {
    let text = try String(contentsOf: fileURL, encoding: .utf8)

    for line in text.split(separator: "\n") {
        let input = line.split(separator: " ")
        let (opponentHand, gameState) = (
            Hand(
                rawValue: String(
                    input
                        .first!))!,
            GameState(
                rawValue: String(
                    input
                        .last!))!
        )

        let roundPoints = gameState.correspondingPoints
        totalPoints +=
            roundPoints + gameState.handPairFor(opponentHand: opponentHand).correspondingPoints
    }
} catch { print("Error: \(fileURL) \n \(error)") }

print("Total points for part 2: \(totalPoints)")
