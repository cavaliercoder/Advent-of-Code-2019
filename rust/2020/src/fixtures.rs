use std::fs;
use std::str::FromStr;

use crate::ToDoError;

pub struct Fixture {
  pub name: String,
  pub data: Vec<u8>,

  cursor: usize,
}

impl Fixture {
  pub fn open(name: &str) -> Fixture {
    let path = format!("../../inputs/2020/{}.txt", name);
    Fixture {
      name: name.to_string(),
      data: fs::read(path).unwrap(),
      cursor: 0,
    }
  }

  /// Parse each line of a fixture as T.
  pub fn parse<T>(&mut self) -> Result<Vec<T>, ToDoError>
  where
    T: FromStr
  {
    let mut values = Vec::new();
    for line in self {
      values.push(line.parse::<T>().map_err(|_| ToDoError)?);
    }
    Ok(values)
  }
}

/// Read a fixture line by line.
impl Iterator for Fixture {
  type Item = String;

  fn next(&mut self) -> Option<Self::Item> {
    if self.cursor >= self.data.len() {
      return None;
    }
    let mut i = self.cursor;
    let mut buf  = String::new();
    while i < self.data.len() && self.data[i] != b'\n' {
      buf.push(self.data[i] as char);
      i += 1;
    }
    self.cursor = i + 1;
    Some(buf)
  }
}
