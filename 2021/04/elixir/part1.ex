defmodule Bingo do
  def find_first_winner(input) do
    [sequence_line | rest] = input |> String.split("\n", trim: true)

    sequence = sequence_line
      |> String.split(",", trim: true)
      |> Enum.map(&String.to_integer/1)

    boards = rest
      |> Enum.map(fn e -> String.split(e, " ", trim: true) end)
      |> Enum.map(fn e ->
        Enum.map(e, fn n -> {String.to_integer(n), false} end)
      end)
      |> Enum.chunk_every(5)

    Enum.reduce_while(sequence, boards, fn number, boards ->
      updated_boards = boards
        |> Enum.map(fn board -> mark_board(board, number) end)

      finished_boards = updated_boards
        |> Enum.filter(fn board -> has_finished_row(board) || has_finished_column(board) end)

      if Enum.empty?(finished_boards) do
        {:cont, updated_boards}
      else
        board = finished_boards |> Enum.at(0)
        {:halt, {number, board}}
      end
    end)
  end

  def mark_board(board, number) do
    board
    |> Enum.map(fn row ->
      Enum.map(row, fn {value, marked} ->
        cond do
          marked -> {value, true}
          value == number -> {value, true}
          true -> {value, false}
        end
      end)
    end)
  end

  def has_finished_row(board) do
    finished_rows = board
      |> Enum.map(fn row ->
        Enum.map(row, fn {_value, marked} -> marked end)
      end)
      |> Enum.map(fn row -> Enum.all?(row) end)
      |> Enum.filter(fn e -> e end)
      |> Enum.count()

    if finished_rows > 0, do: true, else: false
  end

  def has_finished_column(board) do
    size = length(board)

    marked_counts = for i <- 0..size-1 do
      board
        |> Enum.map(fn row ->
          Enum.map(row, fn {_value, marked} -> marked end)
        end)
        |> Enum.map(fn row -> List.to_tuple(row) end)
        |> Enum.count(fn row -> elem(row, i) end)
    end

    marked_counts
      |> Enum.filter(fn count -> count == size end)
      |> then(fn l -> !Enum.empty?(l) end)
  end
end

{number, winner} = Bingo.find_first_winner(File.read!("../input.txt"))

unmarked_sum = winner
  |> Enum.map(fn row ->
    Enum.filter(row, fn {_value, marked} -> !marked end)
  end)
  |> Enum.map(fn row ->
    Enum.map(row, fn {value, _marked} -> value end)
  end)
  |> List.flatten()
  |> Enum.sum()

IO.puts(number * unmarked_sum)
