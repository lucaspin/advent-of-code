defmodule Solution do

  # 1. find pattern for 1 => ab
  # 2. To find pattern for 3, find 5-digit pattern which contains all digits of the pattern for 1 => fbcad
  # 3. To find pattern for 6, find 6-digit pattern which does not contain all digits for pattern for 1 => cdfgeb
  # 4. Find pattern for 4 => eafb
  # 5. To find pattern for 0, find 6-digit pattern which does not contain all digits for pattern for 4 => cagedb
  # 6. Last remaining 6-digit pattern is 9 => cefabd
  # 7. To find pattern for 5, using the pattern for 9, find the only 5-digit pattern which only has one difference => cdfbe
  # 8. The remaining 5-digit is 2 => gcdfa
  def part_two do
    File.read!("../input.txt")
    |> String.split(["\n"])
    |> Enum.map(fn line ->
      [patterns, outputs] = String.split(line, " | ")

      %{
        patterns: sort(String.split(patterns, " ")),
        outputs: sort(String.split(outputs, " "))
      }
    end)
    |> Enum.map(fn %{patterns: patterns, outputs: outputs} ->
      %{digit_map: build_digit_map(patterns), outputs: outputs}
    end)
    |> Enum.map(fn %{digit_map: digit_map, outputs: outputs} ->
      Enum.reduce(outputs, "", fn pattern, output ->
        output <> Integer.to_string(Map.get(digit_map, pattern))
      end)
    end)
    |> Enum.map(&String.to_integer/1)
    |> Enum.sum()
  end

  defp build_digit_map(patterns) do
    digit_map = %{}

    five_digit_patterns = patterns
      |> Enum.filter(fn pattern -> String.length(pattern) == 5 end)

    six_digit_patterns = patterns
      |> Enum.filter(fn pattern -> String.length(pattern) == 6 end)

    pattern_for_one = patterns
      |> Enum.find(fn pattern -> String.length(pattern) == 2 end)
    digit_map = Map.put(digit_map, pattern_for_one, 1)

    pattern_for_three = five_digit_patterns
      |> Enum.find(fn pattern -> contains_all(pattern_for_one, pattern) end)
    digit_map = Map.put(digit_map, pattern_for_three, 3)
    five_digit_patterns = five_digit_patterns |> Enum.filter(fn pattern -> pattern != pattern_for_three end)

    pattern_for_six = six_digit_patterns
      |> Enum.find(fn pattern -> !contains_all(pattern_for_one, pattern) end)
    digit_map = Map.put(digit_map, pattern_for_six, 6)
    six_digit_patterns = six_digit_patterns |> Enum.filter(fn pattern -> pattern != pattern_for_six end)

    pattern_for_four = patterns
      |> Enum.find(fn pattern -> String.length(pattern) == 4 end)
    digit_map = Map.put(digit_map, pattern_for_four, 4)

    pattern_for_zero = six_digit_patterns
      |> Enum.find(fn pattern -> !contains_all(pattern_for_four, pattern) end)
    digit_map = Map.put(digit_map, pattern_for_zero, 0)
    six_digit_patterns = six_digit_patterns |> Enum.filter(fn pattern -> pattern != pattern_for_zero end)

    [pattern_for_nine] = six_digit_patterns
    digit_map = Map.put(digit_map, pattern_for_nine, 9)

    pattern_for_five = five_digit_patterns
      |> Enum.find(fn pattern -> diff(pattern_for_nine, pattern) == 1 end)
    digit_map = Map.put(digit_map, pattern_for_five, 5)
    five_digit_patterns = five_digit_patterns |> Enum.filter(fn pattern -> pattern != pattern_for_five end)

    [pattern_for_two] = five_digit_patterns
    digit_map = Map.put(digit_map, pattern_for_two, 2)

    pattern_for_seven = patterns
      |> Enum.find(fn pattern -> String.length(pattern) == 3 end)
    digit_map = Map.put(digit_map, pattern_for_seven, 7)

    pattern_for_eight = patterns
      |> Enum.find(fn pattern -> String.length(pattern) == 7 end)

    Map.put(digit_map, pattern_for_eight, 8)
  end

  def contains_all(pattern1, pattern2) do
    pattern1_digits = pattern1 |> String.split("")
    pattern2_digits = pattern2 |> String.split("")
    Enum.all?(pattern1_digits, fn digit -> Enum.member?(pattern2_digits, digit) end)
  end

  def diff(pattern1, pattern2) do
    pattern1_digits = pattern1 |> String.split("")
    pattern2_digits = pattern2 |> String.split("")
    pattern1_digits -- pattern2_digits
      |> Enum.count()
  end

  def sort(strings) do
    strings
    |> Enum.map(fn string ->
      string
      |> String.split("", trim: true)
      |> Enum.sort()
      |> Enum.join()
    end)
  end

  def part_one do
    File.read!("../input.txt")
    |> String.split("\n")
    |> Enum.map(fn e -> String.split(e, " | ") end)
    |> Enum.map(fn [_first, second] -> second end)
    |> Enum.map(fn line -> String.split(line, " ") end)
    |> Enum.flat_map(fn e -> e end)
    |> Enum.count(fn e ->
      cond do
        String.length(e) == 2 -> true
        String.length(e) == 3 -> true
        String.length(e) == 4 -> true
        String.length(e) == 7 -> true
        true -> false
      end
    end)
  end
end

Solution.part_two()
|> IO.inspect()
