import _ from 'lodash';

const deepSortObjectByKey = (obj) => {
  if (Array.isArray(obj)) {
    return obj.map(deepSortObjectByKey);
  }

  if (_.isPlainObject(obj)) {
    return Object.fromEntries(
      Object.entries(obj)
        .sort(([key], [otherKey]) => key.localeCompare(otherKey))
        .map(([key, value]) => [key, deepSortObjectByKey(value)])
    );
  }

  return obj;
};

export default deepSortObjectByKey;
