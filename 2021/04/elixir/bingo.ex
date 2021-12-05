defmodule Bingo do
  def init(input) do
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
      # TODO
    end)
  end
end

Bingo.init(File.read!("./input0.txt"))
