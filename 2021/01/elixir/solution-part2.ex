defmodule Solution do
  def run(input) do
    input
    |> String.split("\n")
    |> Enum.map(&String.to_integer/1)
    |> Enum.chunk_every(3, 1, :discard)
    |> Enum.map(&Enum.sum/1)
    |> Enum.chunk_every(2, 1, :discard)
    |> Enum.count(fn [left, right] -> right > left end)
  end
end

IO.inspect(Solution.run(File.read!("../input0.txt")))
