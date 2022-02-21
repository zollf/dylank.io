import faker from 'faker';

faker.seed(1);

export default function Asset() {
  const title = faker.random.word();

  return {
    id: faker.datatype.uuid(),
    title,
    slug: faker.helpers.slugify(title),
    url: faker.image.imageUrl(),
    createdAt: faker.date.soon().toString(),
    updatedAt: faker.date.soon().toString(),
  };
}
