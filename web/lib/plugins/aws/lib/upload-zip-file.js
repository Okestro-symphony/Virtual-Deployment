import path from 'path';
import fs from 'fs';
import crypto from 'crypto';
import utils from '@serverlessinc/sf-core/src/utils.js';
import setS3UploadEncryptionOptions from '../../../aws/set-s3-upload-encryption-options.js';

const { log } = utils;

export default {
  async uploadZipFile({ filename, s3KeyDirname, basename }) {
    const logger = log.get('deploy:upload');
    if (!basename) basename = filename.split(path.sep).pop();

    // TODO refactor to be async (use util function to compute checksum async)
    const data = fs.readFileSync(filename);
    const fileHash = crypto.createHash('sha256').update(data).digest('base64');

    const artifactStream = fs.createReadStream(filename);
    // As AWS SDK request might be postponed (requests are queued)
    // eventual stream error may crash the process (it's thrown as uncaught if not observed).
    // Below lines prevent that
    let streamError;
    artifactStream.on('error', (error) => (streamError = error));

    const key = `${s3KeyDirname}/${basename}`;
    logger.debug('upload to %s/%s', this.bucketName, key);
    let params = {
      Bucket: this.bucketName,
      Key: key,
      Body: artifactStream,
      ContentType: 'application/zip',
      Metadata: {
        filesha256: fileHash,
      },
    };

    const deploymentBucketObject = this.serverless.service.provider.deploymentBucketObject;
    if (deploymentBucketObject) {
      params = setS3UploadEncryptionOptions(params, deploymentBucketObject);
    }

    const response = await this.provider.request('S3', 'upload', params);
    // Interestingly, if request handling was queued, and stream errored (before being consumed by
    // AWS SDK) then SDK call succeeds without actually uploading a file to S3 bucket.
    // Below line ensures that eventual stream error is communicated
    if (streamError) throw streamError;
    return response;
  },
};
