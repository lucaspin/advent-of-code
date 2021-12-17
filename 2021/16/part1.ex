defmodule BITS do
  def versions do
    bits = bits()
    find_versions(0, bits)
  end

  def find_versions(count, bits) when bit_size(bits) < 8, do: count
  def find_versions(count, bits) do
    {version, rest} = version(bits)
    {type, packet} = type(rest)
    if type == 4 do
      {_value, rest} = find_literal_value("", packet)
      find_versions(count + version, rest)
    else
      count + version + find_versions_on_operator(packet)
    end
  end

  def find_literal_value(final, <<1::1, value::4, rest::bitstring>>) do
    find_literal_value(final <> String.pad_leading(Integer.to_string(value, 2), 4, "0"), rest)
  end

  def find_literal_value(final, <<0::1, value::4, rest::bitstring>>) do
    {final <> String.pad_leading(Integer.to_string(value, 2), 4, "0"), rest}
  end

  def find_versions_on_operator(<<0::1, rest::bitstring>>) do
    <<_bit_length::15, sub_packets::bitstring>> = rest
    find_versions(0, sub_packets)
  end

  def find_versions_on_operator(<<1::1, rest::bitstring>>) do
    <<_num_packets::11, sub_packets::bitstring>> = rest
    find_versions(0, sub_packets)
  end

  def version(bits) do
    <<version::3, rest::bitstring>> = bits
    {version, rest}
  end

  def type(bits) do
    <<type::3, rest::bitstring>> = bits
    {type, rest}
  end

  def bits do
    File.read!("./input.txt")
    |> Base.decode16!()
  end
end

BITS.versions()
|> IO.inspect()
